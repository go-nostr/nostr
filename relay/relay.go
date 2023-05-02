package relay

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/go-nostr/nostr"
	"github.com/go-nostr/nostr/message"
	"nhooyr.io/websocket"
)

// New creates and initializes a new Relay instance with the given Options.
func New(opt *Options) *Relay {
	rl := &Relay{
		Options:      opt,
		conn:         make(map[*websocket.Conn]struct{}),
		messHandlers: make(map[string]nostr.MessageHandler),
		serveMux:     new(http.ServeMux),
	}
	rl.serveMux.HandleFunc("/.well-known/nostr.json", rl.internetIdentifierHandlerFunc)
	rl.serveMux.HandleFunc("/", rl.getConnectionHandlerFunc)
	return rl
}

// Options holds the configuration options for a Relay instance.
type Options struct {
	Name          string             `json:"name,omitempty"`
	Description   string             `json:"description,omitempty"`
	PubKey        string             `json:"pub_key,omitempty"`
	Contact       string             `json:"contact,omitempty"`
	SupportedNIPs []int              `json:"supported_nips,omitempty"`
	Software      string             `json:"software,omitempty"`
	Version       string             `json:"version,omitempty"`
	Limitations   *nostr.Limitations `json:"limitations,omitempty"`
}

// Relay represents a websocket relay server with a set of options and handlers.
type Relay struct {
	*Options

	messHandlers map[string]nostr.MessageHandler
	names        map[string]string
	serveMux     *http.ServeMux
	conn         map[*websocket.Conn]struct{}
	mu           sync.Mutex
}

// Handle registers the handler for the given pattern.
func (rl *Relay) Handle(pattern string, handler http.Handler) {
	rl.serveMux.Handle(pattern, handler)
}

// HandleFunc registers the handler function for the given pattern.
func (rl *Relay) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	rl.serveMux.HandleFunc(pattern, handler)
}

// HandleMessage registers the message handler for the given message type.
func (rl *Relay) HandleMessage(typ string, handler nostr.MessageHandler) {
	rl.messHandlers[typ] = handler
}

// HandleMessageFunc registers the message handler function for the given message type.
func (rl *Relay) HandleMessageFunc(typ string, handler func(mess nostr.Message)) {
	rl.messHandlers[typ] = nostr.MessageHandlerFunc(handler)
}

// Publish broadcasts the given message to all connected clients.
func (rl *Relay) Publish(mess nostr.Message) error {
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

// internetIdentifierHandlerFunc handles the /.well-known/nostr.json route.
// It serves the public key of the relay associated with the provided name query.
func (rl *Relay) internetIdentifierHandlerFunc(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"%v\":\"%v\"}", name, rl.names[name])))
}

// getConnectionHandlerFunc handles the root route ("/") and upgrades the
// HTTP request to a websocket connection.
func (rl *Relay) getConnectionHandlerFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rl.addConn(conn)
	go rl.listenConn(conn)
}

// listenConn listens for messages on the provided websocket connection,
// decodes them, and dispatches them to the appropriate message handler.
func (rl *Relay) listenConn(conn *websocket.Conn) {
	ctx := context.Background()
	defer rl.removeConn(conn)
	for {
		_, r, err := conn.Reader(ctx)
		if err != nil {
			return
		}
		// TODO: add websocket mess. type handling
		var mess message.Message
		if err := json.NewDecoder(r).Decode(&mess); err != nil {
			return
		}
		if rl.messHandlers[string(mess.Type())] == nil {
			return
		}
		go rl.messHandlers[string(mess.Type())].Handle(&mess)
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
