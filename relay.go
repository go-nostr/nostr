package nostr

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

// New TBD
func NewRelay() *Relay {
	rl := &Relay{
		conn:         make(map[*websocket.Conn]struct{}),
		err:          make(chan error),
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

	err          chan error
	messHandlers map[MessageType]MessageHandler
	serveMux     *http.ServeMux
	conn         map[*websocket.Conn]struct{}
	mu           sync.Mutex
}

func (rl *Relay) Handle(t MessageType, h MessageHandler) {
	rl.messHandlers[t] = h
}

func (rl *Relay) HandleFunc(t MessageType, f func(mess Message)) {
	rl.messHandlers[t] = MessageHandlerFunc(f)
}

func (rl *Relay) Publish(m Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	byt, err := m.Marshal()
	if err != nil {
		return err
	}
	rl.mu.Lock()
	defer rl.mu.Unlock()
	for conn := range rl.conn {
		go conn.Write(ctx, websocket.MessageText, byt)
	}
	return nil
}

func (rl *Relay) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rl.serveMux.ServeHTTP(w, r)
}

func (rl *Relay) addConnection(c *websocket.Conn) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	rl.conn[c] = struct{}{}
}

func (rl *Relay) listenConnection(c *websocket.Conn) {
	defer rl.removeConnection(c)
	for {
		_, byt, err := c.Read(context.Background())
		if err != nil {
			rl.err <- err
			return
		}
		var args []json.RawMessage
		if err := json.Unmarshal(byt, &args); err != nil {
			rl.err <- err
			return
		}
		var t MessageType
		if err := json.Unmarshal(args[0], &t); err != nil {
			rl.err <- err
			return
		}
		switch t {
		case MessageTypeRequest:
			var m RequestMessage
			if err := m.Unmarshal(byt); err != nil {
				rl.err <- err
				return
			}
			go rl.messHandlers[MessageTypeRequest].Handle(&m)
		}
	}
}

func (rl *Relay) removeConnection(c *websocket.Conn) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	delete(rl.conn, c)
	c.Close(websocket.StatusNormalClosure, "Info: closing connection")
}

func (rl *Relay) defaultInternetIdentifierHandleFunc(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	// HACK: replace with call to repository
	byt, _ := json.Marshal(struct {
		Names  map[string]string `json:"names,omitempty"`
		Relays []string          `json:"relays,omitempty"`
	}{
		Names: map[string]string{
			name: name,
		},
	})
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Write(byt)
}

func (rl *Relay) defaultSubscribeHandleFunc(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	rl.addConnection(c)
	go rl.listenConnection(c)
}
