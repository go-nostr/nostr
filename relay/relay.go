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
		Options: opt,
		cxn:     make(map[*websocket.Conn]struct{}),
		mux:     new(http.ServeMux),
	}
	rl.mux.HandleFunc("/.well-known/nostr.json", rl.getInternetIdentifierHandlerFunc)
	rl.mux.HandleFunc("/", rl.getRootHandlerFunc)
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

	cxn       map[*websocket.Conn]struct{}
	err       chan error
	errFn     func(err error)
	infoDocFn func() (*InformationDocument, error)
	intIdntFn func(name string) (*InternetIdentifier, error)
	msg       chan message.Message
	msgFn     func(msg message.Message)
	mu        sync.Mutex
	mux       *http.ServeMux
}

// HandleError registers the handler for the given pattern.
func (rl *Relay) HandleError(fn func(err error)) {
	rl.errFn = fn
}

// HandleInternetIdentifierFunc TBD
func (rl *Relay) HandleInternetIdentifierFunc(fn func(name string) (*InternetIdentifier, error)) {
	rl.intIdntFn = fn
}

// HandleMessage registers the message handler function for the given message type.
func (rl *Relay) HandleMessage(fn func(msg message.Message)) {
	rl.msgFn = fn
}

// Publish broadcasts the given message to all connected clients.
func (rl *Relay) Publish(msg message.Message) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	data, err := msg.Marshal()
	if err != nil {
		rl.err <- err
	}
	rl.mu.Lock()
	defer rl.mu.Unlock()
	for c := range rl.cxn {
		go c.Write(ctx, websocket.MessageText, data)
	}
}

// ServeHTTP serves the given HTTP request using the internal ServeMux.
func (rl *Relay) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rl.mux.ServeHTTP(w, r)
}

// addConnection adds the given websocket connection to the set of active connections.
func (rl *Relay) addConnection(cxn *websocket.Conn) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	rl.cxn[cxn] = struct{}{}
}

func (rl *Relay) getInformationDocumentHandlerFunc(w http.ResponseWriter, r *http.Request) {
	infoDoc, err := rl.infoDocFn()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(infoDoc)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

// getInternetIdentifierHandlerFunc handles the /.well-known/nostr.json route.
// It serves the public key of the relay associated with the provided name query.
func (rl *Relay) getInternetIdentifierHandlerFunc(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.WriteHeader(http.StatusOK)
	intIdnt, err := rl.intIdntFn(name)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	data, err := json.Marshal(intIdnt)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

// acceptWebsocketHandlerFunc upgrades the HTTP request to websocket connection
func (rl *Relay) acceptWebsocketHandlerFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	rl.addConnection(conn)
	go rl.listenConnection(context.Background(), conn)
}

// getRootHandlerFunc handles the root route ("/") and upgrades the
// HTTP request to a websocket connection.
func (rl *Relay) getRootHandlerFunc(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") == "application/nostr+json" {
		rl.getInformationDocumentHandlerFunc(w, r)
		return
	}
	rl.acceptWebsocketHandlerFunc(w, r)
}

// listenConnection listens for messages on the provided websocket connection,
// decodes them, and dispatches them to the appropriate message handler.
func (rl *Relay) listenConnection(ctx context.Context, conn *websocket.Conn) {
	defer rl.removeConnection(conn)
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
		var msg message.Message
		if err := json.NewDecoder(r).Decode(&msg); err != nil {
			rl.err <- err
			return
		}
		rl.msg <- msg
		select {
		case err := <-rl.err:
			go rl.errFn(err)
		case msg := <-rl.msg:
			go rl.msgFn(msg)
		case <-ctx.Done():
			return
		}
	}
}

// removeConnection removes the given websocket connection from the set of active connections
// and closes the connection with a normal closure status.
func (rl *Relay) removeConnection(cxn *websocket.Conn) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	delete(rl.cxn, cxn)
	cxn.Close(websocket.StatusNormalClosure, "closing connection")
}
