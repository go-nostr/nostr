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

		err:   make(chan error),
		conns: make(map[*websocket.Conn]struct{}),
	}
}

// Options TODO
type Options struct{}

// Client TODO
type Client struct {
	*Options

	conns       map[*websocket.Conn]struct{}
	err         chan error
	messHandler nostr.MessageHandler
	mu          sync.Mutex
}

// HandleMessage TODO
func (cl *Client) HandleMessage(handler nostr.MessageHandler) {
	cl.messHandler = handler
}

// HandleMessageFunc TODO
func (cl *Client) HandleMessageFunc(handler func(mess nostr.Message)) {
	cl.messHandler = nostr.MessageHandlerFunc(handler)
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
		if cl.messHandler == nil {
			fmt.Printf("no handler configured\n")
			return
		}
		go cl.messHandler.Handle(&mess)
	}
}

// removeConn TODO
func (cl *Client) removeConn(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	delete(cl.conns, conn)
	conn.Close(websocket.StatusNormalClosure, "closing connection")
}
