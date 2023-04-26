package nostr

import (
	"context"
	"encoding/json"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

// NewClient TBD
func NewClient() *Client {
	return &Client{
		err:   make(chan error),
		conns: make(map[*websocket.Conn]struct{}),
		mess:  make(chan []byte),
	}
}

// Client TBD
type Client struct {
	conns        map[*websocket.Conn]struct{}
	err          chan error
	mess         chan []byte
	messHandlers map[MessageType]MessageHandler
	mu           sync.Mutex
}

func (cl *Client) Handle(typ MessageType, handler MessageHandler) {
	cl.messHandlers[typ] = handler
}

func (cl *Client) HandleFunc(typ MessageType, handler func(mess Message)) {
	cl.messHandlers[typ] = MessageHandlerFunc(handler)
}

// Publish TBD
func (cl *Client) Publish(mess Message) error {
	ctx := context.TODO()
	data, err := mess.Marshal()
	if err != nil {
		return err
	}
	cl.mu.Lock()
	defer cl.mu.Unlock()
	for conn := range cl.conns {
		go conn.Write(ctx, websocket.MessageText, data)
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
	defer cl.removeConn(conn)
	for {
		_, data, err := conn.Read(context.Background())
		if err != nil {
			cl.err <- err
			return
		}
		var args []json.RawMessage
		if err := json.Unmarshal(data, &args); err != nil {
			cl.err <- err
			return
		}
		var typ MessageType
		if err := json.Unmarshal(args[0], &typ); err != nil {
			cl.err <- err
			return
		}
		switch typ {
		case MessageTypeAuth:
			var mess AuthMessage
			if err := mess.Unmarshal(data); err != nil {
				cl.err <- err
				return
			}
			go cl.messHandlers[MessageTypeAuth].Handle(&mess)
		case MessageTypeCount:
			var mess CountMessage
			if err := mess.Unmarshal(data); err != nil {
				cl.err <- err
				return
			}
			go cl.messHandlers[MessageTypeCount].Handle(&mess)
		case MessageTypeEOSE:
			var mess EOSEMessage
			if err := mess.Unmarshal(data); err != nil {
				cl.err <- err
				return
			}
			go cl.messHandlers[MessageTypeEOSE].Handle(&mess)
		case MessageTypeEvent:
			var mess EventMessage
			if err := mess.Unmarshal(data); err != nil {
				cl.err <- err
				return
			}
			go cl.messHandlers[MessageTypeEvent].Handle(&mess)
		case MessageTypeNotice:
			var mess NoticeMessage
			if err := mess.Unmarshal(data); err != nil {
				cl.err <- err
				return
			}
			go cl.messHandlers[MessageTypeNotice].Handle(&mess)
		case MessageTypeOk:
			var mess OkMessage
			if err := mess.Unmarshal(data); err != nil {
				cl.err <- err
				return
			}
			go cl.messHandlers[MessageTypeOk].Handle(&mess)
		}
	}
}

// removeConn TBD
func (cl *Client) removeConn(conn *websocket.Conn) {
	cl.mu.Lock()
	defer cl.mu.Unlock()
	delete(cl.conns, conn)
	conn.Close(websocket.StatusNormalClosure, "Info: closing connection")
}
