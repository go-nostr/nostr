package client

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/go-nostr/nostr"
	"github.com/go-nostr/nostr/message"
	"nhooyr.io/websocket"
)

// New TODO
func New(opt *Options) *Client {
	return &Client{
		Options: opt,

		err:          make(chan error),
		conns:        make(map[*websocket.Conn]struct{}),
		messHandlers: map[string]nostr.MessageHandler{},
	}
}

// Options TODO
type Options struct{}

// Client TODO
type Client struct {
	*Options

	conns        map[*websocket.Conn]struct{}
	err          chan error
	messHandlers map[string]nostr.MessageHandler
	mu           sync.Mutex
}

// HandleMessage TODO
func (cl *Client) HandleMessage(typ string, handler nostr.MessageHandler) {
	cl.messHandlers[typ] = handler
}

// HandleMessageFunc TODO
func (cl *Client) HandleMessageFunc(typ string, handler func(mess nostr.Message)) {
	cl.messHandlers[typ] = nostr.MessageHandlerFunc(handler)
}

// Publish TODO
func (cl *Client) Publish(mess nostr.Message) error {
	ctx := context.Background()
	data, err := mess.Marshal()
	if err != nil {
		return err
	}
	cl.mu.Lock()
	defer cl.mu.Unlock()
	for conn := range cl.conns {
		if err := conn.Write(ctx, websocket.MessageText, data); err != nil {
			return err
		}
	}
	return nil
}

// Subscribe TODO
func (cl *Client) Subscribe(ctx context.Context, u string) error {
	conn, _, err := websocket.Dial(ctx, u, &websocket.DialOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		return err
	}
	conn.SetReadLimit(6.4e+7)
	if err != nil {
		return err
	}
	cl.addConn(conn)
	go cl.listenConn(conn)
	return nil
}

// addConn TODO
func (cl *Client) addConn(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.conns[conn] = struct{}{}
}

// listenConn TODO
func (cl *Client) listenConn(conn *websocket.Conn) {
	ctx := context.Background()
	defer cl.removeConn(conn)
	for {
		_, r, err := conn.Reader(ctx)
		if err != nil {
			cl.err <- err
			return
		}
		// TODO: add websocket mess. type handling
		var mess message.Message
		if err := json.NewDecoder(r).Decode(&mess); err != nil {
			fmt.Println(err.Error())
			cl.err <- err
			return
		}
		typ := mess.Type()
		if cl.messHandlers[string(typ)] == nil {
			fmt.Printf("no handler configured for message type: %s\n", typ)
			return
		}
		go cl.messHandlers[string(typ)].Handle(&mess)
	}
}

// removeConn TODO
func (cl *Client) removeConn(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	delete(cl.conns, conn)
	conn.Close(websocket.StatusNormalClosure, "closing connection")
}
