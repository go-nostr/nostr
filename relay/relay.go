package relay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-nostr/nostr/message"
	"nhooyr.io/websocket"
)

// New creates and initializes a new Relay instance with the given Options.
func New(opt *Options) *Relay {
	rl := &Relay{
		Options:  opt,
		conn:     make(map[*websocket.Conn]struct{}),
		serveMux: new(http.ServeMux),
	}
	rl.serveMux.HandleFunc("/.well-known/nostr.json", rl.getInternetIdentifierHandlerFunc)
	rl.serveMux.HandleFunc("/", rl.getRootHandlerFunc)
	return rl
}

// Options holds the configuration options for a Relay instance.
type Options struct {
	Name          string
	Description   string
	PubKey        string
	Contact       string
	SupportedNIPs []int
	Software      string
	Version       string
	Limitations   *Limitations
}

// Relay represents a websocket relay server with a set of options and handlers.
type Relay struct {
	*Options

	err                            chan error
	errHandlerFunc                 func(err error)
	internetIdentifierHandlerFunc  func(name string) (*InternetIdentifier, error)
	mess                           chan *message.Message
	messHandler                    message.Handler
	informationDocumentHandlerFunc func() (*InformationDocument, error)
	serveMux                       *http.ServeMux
	conn                           map[*websocket.Conn]struct{}
	mu                             sync.Mutex
}

// Handle registers the handler for the given pattern.
func (rl *Relay) Handle(pattern string, handler http.Handler) {
	rl.serveMux.Handle(pattern, handler)
}

// HandleError registers the handler for the given pattern.
func (rl *Relay) HandleErrorFunc(handler func(err error)) {
	rl.errHandlerFunc = handler
}

// HandleFunc registers the handler function for the given pattern.
func (rl *Relay) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	rl.serveMux.HandleFunc(pattern, handler)
}

// HandleInternetIdentifierFunc TBD
func (rl *Relay) HandleInternetIdentifierFunc(handler func(name string) (*InternetIdentifier, error)) {
	rl.internetIdentifierHandlerFunc = handler
}

// HandleMessage registers the message handler for the given message type.
func (rl *Relay) HandleMessage(handler message.Handler) {
	rl.messHandler = handler
}

// HandleMessageFunc registers the message handler function for the given message type.
func (rl *Relay) HandleMessageFunc(handler func(mess *message.Message)) {
	rl.messHandler = message.HandlerFunc(handler)
}

// Publish broadcasts the given message to all connected clients.
func (rl *Relay) Publish(mess *message.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	data, err := mess.Marshal()
	if err != nil {
		return err
	}
	rl.mu.Lock()
	defer rl.mu.Unlock()
	for conn := range rl.conn {
		go conn.Write(ctx, websocket.MessageText, data)
	}
	return nil
}

// ServeHTTP serves the given HTTP request using the internal ServeMux.
func (rl *Relay) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rl.serveMux.ServeHTTP(w, r)
}

// addConn adds the given websocket connection to the set of active connections.
func (rl *Relay) addConn(conn *websocket.Conn) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	rl.conn[conn] = struct{}{}
}

// getInternetIdentifierHandlerFunc handles the /.well-known/nostr.json route.
// It serves the public key of the relay associated with the provided name query.
func (rl *Relay) getInternetIdentifierHandlerFunc(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	ii, err := rl.internetIdentifierHandlerFunc(name)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	data, err := json.Marshal(ii)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Write(data)
}

// getRootHandlerFunc handles the root route ("/") and upgrades the
// HTTP request to a websocket connection.
func (rl *Relay) getRootHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") == "application/nostr+json" {
		infoDoc, err := rl.informationDocumentHandlerFunc()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		data, err := json.Marshal(infoDoc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
	}
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	rl.addConn(conn)
	go rl.listenConn(context.Background(), conn)
}

// listenConn listens for messages on the provided websocket connection,
// decodes them, and dispatches them to the appropriate message handler.
func (rl *Relay) listenConn(ctx context.Context, conn *websocket.Conn) {
	defer rl.removeConn(conn)
	for {
		typ, r, err := conn.Reader(ctx)
		if err != nil {
			rl.err <- err
			return
		}
		if typ != websocket.MessageText {
			rl.err <- fmt.Errorf("invalid message type")
			return
		}
		var mess message.Message
		if err := json.NewDecoder(r).Decode(&mess); err != nil {
			rl.err <- err
			return
		}
		rl.mess <- &mess
		select {
		case err := <-rl.err:
			go rl.errHandlerFunc(err)
		case mess := <-rl.mess:
			go rl.messHandler.Handle(mess)
		case <-ctx.Done():
			return
		}
	}
}

// removeConn removes the given websocket connection from the set of active connections
// and closes the connection with a normal closure status.
func (rl *Relay) removeConn(conn *websocket.Conn) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	delete(rl.conn, conn)
	conn.Close(websocket.StatusNormalClosure, "Info: closing connection")
}
