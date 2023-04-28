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
func NewRelay() *Relay {
	rl := &Relay{
		conn: make(map[*websocket.Conn]struct{}),
		err:  make(chan error),
		errHandler: func(err error) {
			fmt.Printf("Error: %v", err)
		},
		messHandlers: make(map[MessageType]MessageHandler),
		serveMux:     new(http.ServeMux),
	}
	rl.serveMux.HandleFunc("/.well-known/nostr.json", rl.internetIdentifierHandlerFunc)
	rl.serveMux.HandleFunc("/", rl.getConnectionHandlerFunc)
	return rl
}

// Relay TBD
type Relay struct {
	Name          string       `json:"name,omitempty"`
	Description   string       `json:"description,omitempty"`
	PubKey        string       `json:"pub_key,omitempty"`
	Contact       string       `json:"contact,omitempty"`
	SupportedNIPs []NIP        `json:"supported_nips,omitempty"`
	Software      string       `json:"software,omitempty"`
	Version       string       `json:"version,omitempty"`
	Limitations   *Limitations `json:"limitations,omitempty"`

	err           chan error
	errHandler    func(err error)
	eventHandlers map[EventKind]EventHandler
	messHandlers  map[MessageType]MessageHandler
	names         map[string]string
	serveMux      *http.ServeMux
	conn          map[*websocket.Conn]struct{}
	mu            sync.Mutex
}

// Handle TBD
func (rl *Relay) Handle(pattern string, handler http.Handler) {
	rl.serveMux.Handle(pattern, handler)
}

// HandleFunc TBD
func (rl *Relay) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	rl.serveMux.HandleFunc(pattern, handler)
}

// HandleEvent TBD
func (rl *Relay) HandleEvent(kind EventKind, handler EventHandler) {
	rl.eventHandlers[kind] = handler
}

// HandleEventFunc TBD
func (rl *Relay) HandleEventFunc(kind EventKind, handler func(kind EventKind, evt Event)) {
	rl.eventHandlers[kind] = EventHandlerFunc(handler)
}

// HandleMessage TBD
func (rl *Relay) HandleMessage(typ MessageType, handler MessageHandler) {
	rl.messHandlers[typ] = handler
}

// HandleMessageFunc TBD
func (rl *Relay) HandleMessageFunc(typ MessageType, handler func(mess Message)) {
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
		rl.err <- err
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
		_, data, err := conn.Read(ctx)
		if err != nil {
			rl.err <- err
			return
		}
		var args []json.RawMessage
		if err := json.Unmarshal(data, &args); err != nil {
			rl.err <- err
			return
		}
		var typ MessageType
		if err := json.Unmarshal(args[0], &typ); err != nil {
			rl.err <- err
			return
		}
		switch typ {
		case MessageTypeAuth:
			var mess AuthMessage
			mess.Unmarshal(data)
			go rl.messHandlers[typ].Handle(&mess)
		case MessageTypeClose:
			var mess CloseMessage
			mess.Unmarshal(data)
			go rl.messHandlers[typ].Handle(&mess)
		case MessageTypeCount:
			var mess CountMessage
			mess.Unmarshal(data)
			go rl.messHandlers[typ].Handle(&mess)
		case MessageTypeEOSE:
			var mess EOSEMessage
			mess.Unmarshal(data)
			go rl.messHandlers[typ].Handle(&mess)
		case MessageTypeEvent:
			var mess EventMessage
			mess.Unmarshal(data)
			go rl.messHandlers[typ].Handle(&mess)
		case MessageTypeNotice:
			var mess NoticeMessage
			mess.Unmarshal(data)
			go rl.messHandlers[typ].Handle(&mess)
		case MessageTypeOk:
			var mess OkMessage
			mess.Unmarshal(data)
			go rl.messHandlers[typ].Handle(&mess)
		case MessageTypeRequest:
			var mess RequestMessage
			mess.Unmarshal(data)
			go rl.messHandlers[typ].Handle(&mess)
		default:
			data, err := NewNoticeMessage("unrecognized message type").Marshal()
			if err != nil {
				rl.err <- err
				return
			}
			conn.Write(ctx, websocket.MessageText, data)
		}
	}
}

// removeConn TBD
func (rl *Relay) removeConn(conn *websocket.Conn) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	delete(rl.conn, conn)
	conn.Close(websocket.StatusNormalClosure, "Info: closing connection")
}
