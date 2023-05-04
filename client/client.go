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

		err:   make(chan error),
		conns: make(map[*websocket.Conn]struct{}),
	}
}

// Options TODO
type Options struct{}

// Client TODO
type Client struct {
	*Options

	conns          map[*websocket.Conn]struct{}
	err            chan error
	errHandlerFunc func(err error)
	mess           chan *message.Message
	messHandler    message.Handler
	mu             sync.Mutex
}

// HandleError TODO
func (cl *Client) HandleErrorFunc(handler func(err error)) {
	cl.errHandlerFunc = handler
}

// HandleMessage TODO
func (cl *Client) HandleMessage(handler message.Handler) {
	cl.messHandler = handler
}

// HandleMessageFunc TODO
func (cl *Client) HandleMessageFunc(handler func(mess *message.Message)) {
	cl.messHandler = message.HandlerFunc(handler)
}

func (cl *Client) Listen(ctx context.Context) error {
	for {
		select {
		case err := <-cl.err:
			go cl.errHandlerFunc(err)
		case mess := <-cl.mess:
			go cl.messHandler.Handle(mess)
		case <-ctx.Done():
			return nil
		}
	}
}

// Publish TODO
func (cl *Client) Publish(ctx context.Context, mess *message.Message) error {
	data, err := mess.Marshal()
	if err != nil {
		cl.err <- err
		return err
	}
	cl.mu.Lock()
	defer cl.mu.Unlock()
	for conn := range cl.conns {
		if err := conn.Write(ctx, websocket.MessageText, data); err != nil {
			cl.err <- err
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
		cl.err <- err
		return err
	}
	conn.SetReadLimit(6.4e+7)
	if err != nil {
		cl.err <- err
		return err
	}
	cl.addConnection(conn)
	go cl.listenConnection(ctx, conn)
	return nil
}

// addConnection TODO
func (cl *Client) addConnection(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.conns[conn] = struct{}{}
}

// listenConnection TODO
func (cl *Client) listenConnection(ctx context.Context, conn *websocket.Conn) {
	defer cl.removeConnection(conn)
	for {
		_, r, err := conn.Reader(ctx)
		if err != nil {
			cl.err <- err
			return
		}
		// TODO: add websocket mess. type handling
		var mess message.Message
		if err := json.NewDecoder(r).Decode(&mess); err != nil {
			cl.err <- err
			return
		}
		cl.mess <- &mess
	}
}

// removeConnection TODO
func (cl *Client) removeConnection(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	delete(cl.conns, conn)
	if err := conn.Close(websocket.StatusNormalClosure, "closing connection"); err != nil {
		cl.err <- err
	}
}
