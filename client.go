package nostr

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"nhooyr.io/websocket"
)

// NewClient TBD
func NewClient(opt *ClientOptions) *Client {
	return &Client{
		ClientOptions: opt,

		err:          make(chan error),
		conns:        make(map[*websocket.Conn]struct{}),
		messHandlers: map[string]MessageHandler{},
	}
}

// ClientOptions TBD
type ClientOptions struct{}

// Client TBD
type Client struct {
	*ClientOptions

	conns        map[*websocket.Conn]struct{}
	err          chan error
	messHandlers map[string]MessageHandler
	mu           sync.Mutex
}

// HandleMessage TBD
func (cl *Client) HandleMessage(typ string, handler MessageHandler) {
	cl.messHandlers[typ] = handler
}

// HandleMessageFunc TBD
func (cl *Client) HandleMessageFunc(typ string, handler func(mess Message)) {
	cl.messHandlers[typ] = MessageHandlerFunc(handler)
}

// Publish TBD
func (cl *Client) Publish(mess Message) error {
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

// Subscribe TBD
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

// addConn TBD
func (cl *Client) addConn(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.conns[conn] = struct{}{}
}

// listenConn TBD
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
		var mess RawMessage
		if err := json.NewDecoder(r).Decode(&mess); err != nil {
			fmt.Printf("Error: %v", err.Error())
			cl.err <- err
			return
		}
		typ := mess.Type()
		if cl.messHandlers[string(typ)] == nil {
			fmt.Printf("Warn: no handler configured for message type: %v\n", typ)
			return
		}
		go cl.messHandlers[string(typ)].Handle(&mess)
	}
}

// removeConn TBD
func (cl *Client) removeConn(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	delete(cl.conns, conn)
	conn.Close(websocket.StatusNormalClosure, "Info: closing connection")
}
