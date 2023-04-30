package nostr

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

// New TBD
func NewRelay(opt *RelayOptions) *Relay {
	rl := &Relay{
		RelayOptions: opt,
		conn:         make(map[*websocket.Conn]struct{}),
		messHandlers: make(map[string]MessageHandler),
		serveMux:     new(http.ServeMux),
	}
	rl.serveMux.HandleFunc("/.well-known/nostr.json", rl.internetIdentifierHandlerFunc)
	rl.serveMux.HandleFunc("/", rl.getConnectionHandlerFunc)
	return rl
}

type RelayOptions struct {
	Name          string       `json:"name,omitempty"`
	Description   string       `json:"description,omitempty"`
	PubKey        string       `json:"pub_key,omitempty"`
	Contact       string       `json:"contact,omitempty"`
	SupportedNIPs []int        `json:"supported_nips,omitempty"`
	Software      string       `json:"software,omitempty"`
	Version       string       `json:"version,omitempty"`
	Limitations   *Limitations `json:"limitations,omitempty"`
}

// Relay TBD
type Relay struct {
	*RelayOptions

	messHandlers map[string]MessageHandler
	names        map[string]string
	serveMux     *http.ServeMux
	conn         map[*websocket.Conn]struct{}
	mu           sync.Mutex
}

// Handle TBD
func (rl *Relay) Handle(pattern string, handler http.Handler) {
	rl.serveMux.Handle(pattern, handler)
}

// HandleFunc TBD
func (rl *Relay) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	rl.serveMux.HandleFunc(pattern, handler)
}

// HandleMessage TBD
func (rl *Relay) HandleMessage(typ string, handler MessageHandler) {
	rl.messHandlers[typ] = handler
}

// HandleMessageFunc TBD
func (rl *Relay) HandleMessageFunc(typ string, handler func(mess Message)) {
	rl.messHandlers[typ] = MessageHandlerFunc(handler)
}

// Publish TBD
func (rl *Relay) Publish(mess Message) error {
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

// ServeHTTP TBD
func (rl *Relay) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rl.serveMux.ServeHTTP(w, r)
}

// addConn TBD
func (rl *Relay) addConn(conn *websocket.Conn) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	rl.conn[conn] = struct{}{}
}

// internetIdentifierHandlerFunc TBD
func (rl *Relay) internetIdentifierHandlerFunc(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("{\"%v\":\"%v\"}", name, rl.names[name])))
}

// getConnectionHandlerFunc TBD
func (rl *Relay) getConnectionHandlerFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	rl.addConn(conn)
	go rl.listenConn(conn)
}

// listenConn TBD
func (rl *Relay) listenConn(conn *websocket.Conn) {
	ctx := context.Background()
	defer rl.removeConn(conn)
	for {
		_, r, err := conn.Reader(ctx)
		if err != nil {
			return
		}
		// TODO: add websocket mess. type handling
		var mess RawMessage
		if err := json.NewDecoder(r).Decode(&mess); err != nil {
			return
		}
		if rl.messHandlers[string(mess.Type())] == nil {
			return
		}
		go rl.messHandlers[string(mess.Type())].Handle(mess)
	}
}

// removeConn TBD
func (rl *Relay) removeConn(conn *websocket.Conn) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	delete(rl.conn, conn)
	conn.Close(websocket.StatusNormalClosure, "Info: closing connection")
}
