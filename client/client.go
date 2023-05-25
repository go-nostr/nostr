package client

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/go-nostr/nostr/message"
	"nhooyr.io/websocket"
)

// New creates a new Client with the given options.
// It sets up channels for errors and messages, default handlers for errors and messages,
// and a map for active WebSocket connections.
func New(opt *Options) *Client {
	return &Client{
		Options: opt,

		errCh: make(chan error),
		errFn: func(err error) {
			fmt.Printf("No error handler registered")
		},
		msgCh: make(chan message.Message),
		msgFn: func(msg message.Message) {
			fmt.Printf("No message handler registered.")
		},
		connMap: make(map[*websocket.Conn]struct{}),
	}
}

// Options represents the configuration options for a Client.
// It currently only includes a read limit for WebSocket connections.
type Options struct {
	ReadLimit int64
}

// Client is a structure representing a client in a WebSocket communication setup.
// It contains options for configuration, channels for errors and messages,
// handlers for errors and messages, a map for managing active connections,
// and a Mutex for safe concurrent access.
type Client struct {
	*Options

	connMap map[*websocket.Conn]struct{}
	errCh   chan error
	errFn   func(err error)
	msgCh   chan message.Message
	msgFn   func(msg message.Message)
	mu      sync.Mutex
}

// Connect establishes a WebSocket connection to the given URL.
// It adds the new connection to the client's map of connections and starts listening on it.
func (cl *Client) Connect(ctx context.Context, u string) {
	conn, _, err := websocket.Dial(ctx, u, &websocket.DialOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		cl.errCh <- err
	}
	if cl.Options != nil {
		conn.SetReadLimit(cl.ReadLimit)
	}
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.connMap[conn] = struct{}{}
	go cl.listenConnection(ctx, conn)
}

// HandleErrorFunc sets the function to be called when an error occurs.
// The function should take an error as its argument.
func (cl *Client) HandleErrorFunc(fn func(error)) {
	cl.errFn = fn
}

// HandleMessageFunc sets the function to be called when a message is received.
// The function should take a message as its argument.
func (cl *Client) HandleMessageFunc(fn func(message.Message)) {
	cl.msgFn = fn
}

// Listen starts listening for errors and messages on the client's channels.
// When a message or an error is received, it calls the appropriate handler function in a new goroutine.
// The function returns when the provided context is done.
func (cl *Client) Listen(ctx context.Context) error {
	for {
		select {
		case err := <-cl.errCh:
			go cl.errFn(err)
		case msg := <-cl.msgCh:
			go cl.msgFn(msg)
		case <-ctx.Done():
			return nil
		}
	}
}

// SendMessage sends the given message to all active WebSocket connections of the client.
// It first marshals the message into a byte slice. If an error occurs, it sends the error on the error channel.
func (cl *Client) SendMessage(ctx context.Context, msg message.Message) {
	data, err := msg.Marshal()
	if err != nil {
		cl.errCh <- err
	}
	cl.mu.Lock()
	defer cl.mu.Unlock()
	for conn := range cl.connMap {
		if err := conn.Write(ctx, websocket.MessageText, data); err != nil {
			cl.errCh <- err
		}
	}
}

// listenConnection starts listening for messages on a WebSocket connection.
// When a text message is received, it decodes the message and sends it on the message channel.
// If an error occurs or a non-text message is received, it sends the error on the error channel and returns.
func (cl *Client) listenConnection(ctx context.Context, conn *websocket.Conn) {
	defer cl.removeConnection(conn)
	for {
		typ, rdr, err := conn.Reader(ctx)
		if err != nil {
			cl.errCh <- err
			return
		}
		if typ != websocket.MessageText {
			cl.errCh <- err
			return
		}
		var msg message.Message
		if err := json.NewDecoder(rdr).Decode(&msg); err != nil {
			cl.errCh <- err
			return
		}
		select {
		case <-ctx.Done():
			return
		default:
			cl.msgCh <- msg
		}
	}
}

// removeConnection removes a WebSocket connection from the client's map of connections and closes it.
// If an error occurs while closing the connection, it sends the error on the error channel.
func (cl *Client) removeConnection(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	delete(cl.connMap, conn)
	if err := conn.Close(websocket.StatusNormalClosure, "closing connection"); err != nil {
		cl.errCh <- err
	}
}
