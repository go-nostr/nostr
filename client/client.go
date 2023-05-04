package client

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/go-nostr/nostr/message"
	"nhooyr.io/websocket"
)

// New TODO
func New(opt *Options) *Client {
	return &Client{
		Options: opt,

		err:  make(chan error),
		conn: make(map[*websocket.Conn]struct{}),
	}
}

// Options TODO
type Options struct {
	ReadLimit int64
}

// Client TODO
type Client struct {
	*Options

	conn  map[*websocket.Conn]struct{}
	err   chan error
	errFn func(err error)
	msg   chan message.Message
	msgFn func(msg message.Message)
	mu    sync.Mutex
}

// HandleError TODO
func (cl *Client) HandleErrorFunc(fn func(err error)) {
	cl.errFn = fn
}

// HandleMessage TODO
func (cl *Client) HandleMessageFunc(fn func(msg message.Message)) {
	cl.msgFn = fn
}

// Listen TODO
func (cl *Client) Listen(ctx context.Context) error {
	for {
		select {
		case err := <-cl.err:
			if cl.errFn != nil {
				go cl.errFn(err)
			}
		case msg := <-cl.msg:
			if cl.msgFn != nil {
				go cl.msgFn(msg)
			}
		case <-ctx.Done():
			return nil
		}
	}
}

// Publish TODO
func (cl *Client) Publish(ctx context.Context, msg message.Message) {
	data, err := msg.Marshal()
	if err != nil {
		cl.err <- err
	}
	cl.mu.Lock()
	defer cl.mu.Unlock()
	for conn := range cl.conn {
		if err := conn.Write(ctx, websocket.MessageText, data); err != nil {
			cl.err <- err
		}
	}
}

// Subscribe TODO
func (cl *Client) Subscribe(ctx context.Context, u string) {
	conn, _, err := websocket.Dial(ctx, u, &websocket.DialOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		cl.err <- err
	}
	if cl.Options != nil {
		conn.SetReadLimit(cl.ReadLimit)
	}
	cl.addConn(conn)
	go cl.listenConn(ctx, conn)
}

// addConnection TODO
func (cl *Client) addConn(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.conn[conn] = struct{}{}
}

// listenConnection TODO
func (cl *Client) listenConn(ctx context.Context, conn *websocket.Conn) {
	defer cl.removeConn(conn)
	for {
		typ, rdr, err := conn.Reader(ctx)
		if err != nil {
			go cl.errFn(err)
			return
		}
		if typ != websocket.MessageText {
			go cl.errFn(err)
			return
		}
		var msg message.Message
		if err := json.NewDecoder(rdr).Decode(&msg); err != nil {
			go cl.errFn(err)
			return
		}
		select {
		case <-ctx.Done():
			return
		default:
			go cl.msgFn(msg)
		}
	}
}

// removeConnection TODO
func (cl *Client) removeConn(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	delete(cl.conn, conn)
	if err := conn.Close(websocket.StatusNormalClosure, "closing connection"); err != nil {
		cl.err <- err
	}
}
