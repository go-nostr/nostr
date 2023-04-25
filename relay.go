package nostr

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

// MessageHandlerFunc ...
type MessageHandlerFunc func(mess Message)

// Handle ...
func (f MessageHandlerFunc) Handle(mess Message) {
	f(mess)
}

type MessageHandler interface {
	Handle(mess Message)
}

// New TBD
func NewRelay() *Relay {
	rl := &Relay{
		conns:        make(map[*websocket.Conn]struct{}),
		messHandlers: make(map[MessageType]MessageHandler),
		serveMux:     new(http.ServeMux),
	}
	rl.serveMux.HandleFunc("/.well-known/nostr.json", rl.defaultInternetIdentifierHandleFunc)
	rl.serveMux.HandleFunc("/", rl.defaultSubscribeHandleFunc)
	return rl
}

// Relay TBD
type Relay struct {
	Name          string       `json:"name,omitempty"`
	Description   string       `json:"description,omitempty"`
	PubKey        string       `json:"pub_key,omitempty"`
	Contact       string       `json:"contact,omitempty"`
	SupportedNIPs []string     `json:"supported_nips,omitempty"`
	Software      string       `json:"software,omitempty"`
	Version       string       `json:"version,omitempty"`
	Limitations   *Limitations `json:"limitations,omitempty"`

	messHandlers map[MessageType]MessageHandler
	serveMux     *http.ServeMux
	conns        map[*websocket.Conn]struct{}
	mu           sync.Mutex
	// TODO: add err chan
}

func (rl *Relay) Handle(messType MessageType, messHandler MessageHandler) {
	rl.messHandlers[messType] = messHandler
}

func (rl *Relay) HandleFunc(messType MessageType, h func(mess Message)) {
	rl.messHandlers[messType] = MessageHandlerFunc(h)
}

func (rl *Relay) Publish(mess Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	data, err := mess.Marshal()
	if err != nil {
		return err
	}
	rl.mu.Lock()
	defer rl.mu.Unlock()
	for conn := range rl.conns {
		conn.Write(ctx, websocket.MessageText, data)
	}
	return nil
}

func (rl *Relay) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rl.serveMux.ServeHTTP(w, r)
}

func (rl *Relay) addConn(cl *websocket.Conn) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	rl.conns[cl] = struct{}{}
}

func (rl *Relay) listenConn(conn *websocket.Conn) {
	defer rl.removeConn(conn)
	for {
		_, data, err := conn.Read(context.Background())
		if err != nil {
			fmt.Printf("error reading from connection: %v", err)
			return
		}
		var args []json.RawMessage
		if err := json.Unmarshal(data, &args); err != nil {
			fmt.Printf("error unmarshaling arguments: %v", err)
			return
		}
		var messType MessageType
		if err := json.Unmarshal(args[0], &messType); err != nil {
			fmt.Printf("error unmarshaling message type: %v", err)
			return
		}
		switch messType {
		case MessageTypeRequest:
			var reqMessage RequestMessage
			if err := reqMessage.Unmarshal(data); err != nil {
				fmt.Printf("error unmarshaling request message: %v", err)
				return
			}
			go rl.messHandlers[MessageTypeRequest].Handle(&reqMessage)
		default:
			fmt.Printf("Read: %s", data)
		}
	}
}

func (rl *Relay) removeConn(conn *websocket.Conn) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	delete(rl.conns, conn)
	conn.Close(websocket.StatusNormalClosure, "Info: closing connection")
}

func (rl *Relay) defaultInternetIdentifierHandleFunc(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	// HACK: replace with call to repository
	data, _ := json.Marshal(struct {
		Names  map[string]string `json:"names,omitempty"`
		Relays []string          `json:"relays,omitempty"`
	}{
		Names: map[string]string{
			name: name,
		},
	})
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

func (rl *Relay) defaultSubscribeHandleFunc(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	rl.addConn(conn)
	go rl.listenConn(conn)
}
