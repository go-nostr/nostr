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
	}
}

// Client TBD
type Client struct {
	conns         map[*websocket.Conn]struct{}
	err           chan error
	eventHandlers map[EventKind]EventHandler
	messHandlers  map[MessageType]MessageHandler
	mu            sync.Mutex
}

// HandleEvent TBD
func (cl *Client) HandleEvent(kind EventKind, handler EventHandler) {
	cl.eventHandlers[kind] = handler
}

// HandleEventFunc TBD
func (cl *Client) HandleEventFunc(kind EventKind, handler func(kind EventKind, evt Event)) {
	cl.eventHandlers[kind] = EventHandlerFunc(handler)
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
	ctx := context.Background()
	defer cl.removeConn(conn)
	for {
		_, data, err := conn.Read(ctx)
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
			mess.Unmarshal(data)
			go cl.messHandlers[typ].Handle(&mess)
		case MessageTypeClose:
			var mess CloseMessage
			mess.Unmarshal(data)
			go cl.messHandlers[typ].Handle(&mess)
		case MessageTypeCount:
			var mess CountMessage
			mess.Unmarshal(data)
			go cl.messHandlers[typ].Handle(&mess)
		case MessageTypeEOSE:
			var mess EOSEMessage
			mess.Unmarshal(data)
			go cl.messHandlers[typ].Handle(&mess)
		case MessageTypeEvent:
			var mess EventMessage
			mess.Unmarshal(data)
			go cl.messHandlers[typ].Handle(&mess)
		case MessageTypeNotice:
			var mess NoticeMessage
			mess.Unmarshal(data)
			go cl.messHandlers[typ].Handle(&mess)
		case MessageTypeOk:
			var mess OkMessage
			mess.Unmarshal(data)
			go cl.messHandlers[typ].Handle(&mess)
		case MessageTypeRequest:
			var mess RequestMessage
			mess.Unmarshal(data)
			go cl.messHandlers[typ].Handle(&mess)
		default:
			data, err := NewNoticeMessage("unrecognized message type").Marshal()
			if err != nil {
				cl.err <- err
				return
			}
			conn.Write(ctx, websocket.MessageText, data)
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
