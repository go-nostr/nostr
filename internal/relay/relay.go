package relay

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/go-nostr/nostr"
	"nhooyr.io/websocket"
)

const (
	defaultHostname = "0.0.0.0"
	defaultPort     = 4317
)

type internetIdentifierHandler struct {
}

func (h *internetIdentifierHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, _ := json.Marshal(struct {
		Names  []string `json:"names,omitempty"`
		Relays []string `json:"relays,omitempty"`
	}{
		Names:  []string{"bob"},
		Relays: []string{},
	})
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Content-Type", "application/json")
	w.Write(data)
}

type subscribeHandler struct {
}

func (h *subscribeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "")
	// TODO: add service call to add subscriber
	if errors.Is(err, context.Canceled) {
		return
	}
	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}
	if err != nil {
		log.Printf("%v", err)
		return
	}
}

func NewRelay() *Relay {
	serveMux := &http.ServeMux{}
	serveMux.Handle("/.well-known/nostr.json", &internetIdentifierHandler{})
	serveMux.Handle("/", &subscribeHandler{})
	return &Relay{
		conns: make(map[*websocket.Conn]struct{}),
		server: &http.Server{
			Addr:    fmt.Sprintf("%+v:%+v", defaultHostname, defaultPort),
			Handler: serveMux,
		},
	}
}

type Relay struct {
	Name          string             `json:"name,omitempty"`
	Description   string             `json:"description,omitempty"`
	PubKey        string             `json:"pub_key,omitempty"`
	Contact       string             `json:"contact,omitempty"`
	SupportedNIPs []string           `json:"supported_nips,omitempty"`
	Software      string             `json:"software,omitempty"`
	Version       string             `json:"version,omitempty"`
	Limitations   *nostr.Limitations `json:"limitations,omitempty"`

	messHandlers map[nostr.MessageType]func(mess nostr.Message)
	server       *http.Server
	conns        map[*websocket.Conn]struct{}
	mess         chan []byte
	mu           sync.Mutex
}

func (r *Relay) Handle(typ nostr.MessageType, fn func(mess nostr.Message)) {
	r.messHandlers[typ] = fn
}

func (r *Relay) ListenAndServe() error {
	return r.server.ListenAndServe()
}

func (r *Relay) Serve(l net.Listener) error {
	return r.server.Serve(l)
}

func (r *Relay) Publish(mess nostr.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	data, err := mess.Marshal()
	if err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	for conn := range r.conns {
		conn.Write(ctx, websocket.MessageText, data)
	}
	return nil
}

func (r *Relay) Subscribe(u string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	conn, _, err := websocket.Dial(ctx, u, &websocket.DialOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		return err
	}
	r.addConnection(conn)
	go r.listenConnection(conn)
	return nil
}

func (r *Relay) addConnection(cl *websocket.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.conns[cl] = struct{}{}
}

func (r *Relay) listenConnection(conn *websocket.Conn) {
	defer r.removeConnection(conn)
	for {
		_, mess, err := conn.Read(context.Background())
		if err != nil {
			fmt.Printf("Error reading from connection: %v\n", err)
			return
		}
		r.mess <- mess
	}
}

func (r *Relay) removeConnection(conn *websocket.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.conns, conn)
	conn.Close(websocket.StatusNormalClosure, "closing connection")
}
