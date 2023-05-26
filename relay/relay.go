package relay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-nostr/nostr/message"
	"nhooyr.io/websocket"
)

// New creates a new Relay instance with the provided Options. If no options are provided,
// it will create default options. If the Origin in options is not set, it sets it to "*".
// This function also initializes a map for connections, sets error and message handlers
// to default functions, and sets two HTTP handlers for ".well-known/nostr.json" and "/"
// routes.
func New(opt *Options) *Relay {
	if opt == nil {
		opt = new(Options)
	}
	if opt.Origin == "" {
		opt.Origin = "*"
	}
	rl := &Relay{
		Options: opt,

		connMap: make(map[*websocket.Conn]struct{}),
		errFn: func(err error) {
			fmt.Printf("No error handler registered")
		},
		msgFn: func(msg message.Message) {
			fmt.Printf("No message handler registered.")
		},
		mux: new(http.ServeMux),
	}
	rl.mux.HandleFunc("/.well-known/nostr.json", rl.getInternetIdentifier)
	rl.mux.HandleFunc("/", rl.getIndex)
	return rl
}

// Options holds the configuration options for a Relay instance. This includes the name,
// description, public key, contact, origin, supported NIPs, software, version, and limitations
// for the relay instance.
type Options struct {
	Name          string
	Description   string
	PubKey        string
	Contact       string
	Origin        string
	SupportedNIPs []int
	Software      string
	Version       string
	Limitations   *Limitations
}

// Relay represents a websocket relay server. It holds options, a map of connections, handlers
// for errors and messages, mutex for concurrent access, and a ServeMux for HTTP request routing.
type Relay struct {
	*Options

	connMap               map[*websocket.Conn]struct{}
	errFn                 func(error)
	informationDocumentFn func() (*InformationDocument, error)
	internetIdentiferFn   func(string) (*InternetIdentifier, error)
	msgFn                 func(message.Message)
	mu                    sync.Mutex
	mux                   *http.ServeMux
}

// HandleErrorFunc registers a function that will handle errors. This function is called when an error
// occurs.
func (rl *Relay) HandleErrorFunc(fn func(error)) {
	rl.errFn = fn
}

// HandleInformationDocumentFunc registers a function that will handle the generation of an information document.
// This function is expected to return an InformationDocument instance or an error.
func (rl *Relay) HandleInformationDocumentFunc(fn func() (*InformationDocument, error)) {
	rl.informationDocumentFn = fn
}

// HandleInternetIdentifierFunc registers a function that will handle the internet identifier. This function is expected
// to return an InternetIdentifier instance or an error given a string input.
func (rl *Relay) HandleInternetIdentifierFunc(fn func(string) (*InternetIdentifier, error)) {
	rl.internetIdentiferFn = fn
}

// HandleMessageFunc registers a function that will handle messages. This function is called with a
// Message as a parameter whenever a message is received.
func (rl *Relay) HandleMessageFunc(fn func(message.Message)) {
	rl.msgFn = fn
}

// SendMessage sends the provided Message to all connected clients. If an error occurs while marshalling
// the message, it calls the registered error handler function.
func (rl *Relay) SendMessage(ctx context.Context, msg message.Message) {
	data, err := msg.Marshal()
	if err != nil {
		go rl.errFn(err)
	}
	rl.mu.Lock()
	defer rl.mu.Unlock()
	for conn := range rl.connMap {
		go conn.Write(ctx, websocket.MessageText, data)
	}
}

// ServeHTTP serves the given HTTP request using the internal ServeMux. It allows the Relay instance
// to be used as an HTTP handler.
func (rl *Relay) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rl.mux.ServeHTTP(w, r)
}

// acceptConnection is an internal function that upgrades an HTTP request to a websocket connection.
// If the upgrade is successful, the connection is added to the active connections map and a goroutine
// is started to listen for messages on the connection.
func (rl *Relay) acceptConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	rl.mu.Lock()
	defer rl.mu.Unlock()
	rl.connMap[conn] = struct{}{}
	go rl.listenConnection(context.Background(), conn)
}

// getIndex handles the root route ("/"). If the "Accept" header is "application/nostr+json", it serves
// the information document. Otherwise, it attempts to upgrade the connection to a websocket.
func (rl *Relay) getIndex(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Accept") == "application/nostr+json" {
		rl.getInformationDocument(w, r)
		return
	}
	rl.acceptConnection(w, r)
}

// getInformationDocument is an internal function that handles the generation of an information document.
// The function uses the registered information document function to generate the document, then writes it
// to the HTTP response in JSON format.
func (rl *Relay) getInformationDocument(w http.ResponseWriter, r *http.Request) {
	informationDocument, err := rl.informationDocumentFn()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	data, err := json.Marshal(informationDocument)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Access-Control-Allow-Origin", rl.Origin)
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

// getInternetIdentifier handles the "/.well-known/nostr.json" route. It retrieves the internet identifier
// associated with the name query parameter using the registered internet identifier function, then writes it
// to the HTTP response in JSON format.
func (rl *Relay) getInternetIdentifier(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	internetIdentifier, err := rl.internetIdentiferFn(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	data, err := json.Marshal(internetIdentifier)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Access-Control-Allow-Origin", rl.Origin)
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

// listenConnection is an internal function that listens for messages on a websocket connection. It reads messages
// from the connection, decodes them into Message objects, and dispatches them to the registered message handler function.
// If an error occurs during this process, it calls the registered error handler function.
func (rl *Relay) listenConnection(ctx context.Context, conn *websocket.Conn) {
	defer rl.removeConnection(conn)
	for {
		typ, rdr, err := conn.Reader(ctx)
		if err != nil {
			go rl.errFn(err)
			return
		}
		if typ != websocket.MessageText {
			go rl.errFn(fmt.Errorf("unsupported message type"))
			return
		}
		var msg message.Message
		if err := json.NewDecoder(rdr).Decode(&msg); err != nil {
			go rl.errFn(err)
			return
		}
		select {
		case <-ctx.Done():
			return
		default:
			go rl.msgFn(msg)
		}
	}
}

// removeConnection is an internal function that removes a websocket connection from the active connections map.
// It also closes the connection with a normal closure status.
func (rl *Relay) removeConnection(conn *websocket.Conn) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	delete(rl.connMap, conn)
	conn.Close(websocket.StatusNormalClosure, "closing connection")
}
