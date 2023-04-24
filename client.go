package nostr

import (
	"context"
	"fmt"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

// NewClient TBD
func NewClient() *Client {
	return &Client{
		conns: make(map[*websocket.Conn]struct{}),
		mess:  make(chan []byte),
	}
}

// Client TBD
type Client struct {
	conns map[*websocket.Conn]struct{}
	mess  chan []byte
	mu    sync.Mutex
}

// Publish TBD
func (cl *Client) Publish(mess Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	byt, err := mess.Marshal()
	if err != nil {
		return err
	}
	cl.mu.Lock()
	defer cl.mu.Unlock()
	for conn := range cl.conns {
		conn.Write(ctx, websocket.MessageText, byt)
	}
	return nil
}

// Subscribe TBD
func (cl *Client) Subscribe(u string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	conn, _, err := websocket.Dial(ctx, u, &websocket.DialOptions{
		CompressionMode: websocket.CompressionDisabled,
	})
	if err != nil {
		return err
	}
	cl.addConnection(conn)
	go cl.listenConnection(conn)
	return nil
}

// addConnection TBD
func (cl *Client) addConnection(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	cl.conns[conn] = struct{}{}
}

// listenConnection TBD
func (cl *Client) listenConnection(conn *websocket.Conn) {
	defer cl.removeConnection(conn)
	for {
		_, mess, err := conn.Read(context.Background())
		if err != nil {
			fmt.Printf("Error reading from relay: %v\n", err)
			return
		}
		cl.mess <- mess
	}
}

// removeConnection TBD
func (cl *Client) removeConnection(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	delete(cl.conns, conn)
	conn.Close(websocket.StatusNormalClosure, "closing connection")
}
