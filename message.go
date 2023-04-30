package nostr

import (
	"encoding/json"
)

const (
	// MessageTypeAuth represents an authentication message type.
	MessageTypeAuth = "AUTH"
	// MessageTypeClose represents a Close message type.
	MessageTypeClose = "CLOSE"
	// MessageTypeCount represents a count message type, usually for counting events.
	MessageTypeCount = "COUNT"
	// MessageTypeEndOfStoredEvents represents an End of Stored Events message type.
	MessageTypeEOSE = "EOSE"
	// MessageTypeEvent represents an Event message type.
	MessageTypeEvent = "EVENT"
	// MessageTypeNotice represents a Notice message type, usually for notifications.
	MessageTypeNotice = "NOTICE"
	// MessageTypeOk represents a success confirmation message type.
	MessageTypeOk = "OK"
	// MessageTypeRequest represents a Request message type.
	MessageTypeRequest = "REQ"
)

// MessageHandlerFunc TBD
type MessageHandlerFunc func(mess Message)

// Handle TBD
func (f MessageHandlerFunc) Handle(mess Message) {
	f(mess)
}

// MessageHandler TBD
type MessageHandler interface {
	Handle(mess Message)
}

// Message is an interface for encoding and marshaling messages.
type Message interface {
	Marshal() ([]byte, error)
	Push(v any) error
	Type() []byte
	Unmarshal(data []byte) error
	Values() []any
}

// NewRawMessage
func NewRawMessage(typ string) Message {
	mess := &RawMessage{}
	mess.Push(typ)
	return mess
}

// RawMessage TBD
type RawMessage []json.RawMessage

// Marshal TBD
func (m *RawMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Push TBD
func (m *RawMessage) Push(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	*m = append(*m, data)
	return nil
}

// Type TBD
func (m *RawMessage) Type() []byte {
	if len(*m) < 1 {
		return []byte{}
	}
	var typ string
	if err := json.Unmarshal((*m)[0], &typ); err != nil {
		return []byte{}
	}
	return []byte(typ)
}

// Unmarshal TBD
func (m *RawMessage) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Values TBD
func (m *RawMessage) Values() []any {
	var vals []any
	for _, arg := range *m {
		var val any
		json.Unmarshal(arg, &val)
		vals = append(vals, val)
	}
	return vals
}

func NewAuthMessage(challenge string, event *Event) Message {
	mess := &AuthMessage{&RawMessage{}}
	mess.Push(MessageTypeAuth)
	if challenge != "" {
		mess.Push(challenge)
	}
	if event != nil {
		mess.Push(event)
	}
	return mess
}

// AuthMessage TBD
type AuthMessage struct {
	*RawMessage
}

func NewCloseMessage() Message {
	mess := &CloseMessage{&RawMessage{}}
	mess.Push(MessageTypeClose)
	return mess
}

// CloseMessage TBD
type CloseMessage struct {
	*RawMessage
}

func NewCountMessage() Message {
	mess := &CountMessage{&RawMessage{}}
	mess.Push(MessageTypeCount)
	return mess
}

// CountMessage TBD
type CountMessage struct {
	*RawMessage
}

func NewEOSEMessage() Message {
	mess := &EOSEMessage{&RawMessage{}}
	mess.Push(MessageTypeEOSE)
	return mess
}

// EOSEMessage TBD
type EOSEMessage struct {
	*RawMessage
}

func NewEventMessage() Message {
	mess := &EventMessage{&RawMessage{}}
	mess.Push(MessageTypeEvent)
	return mess
}

// EventMessage TBD
type EventMessage struct {
	*RawMessage
}

func (m *EventMessage) Event() Event {
	if len(*m.RawMessage) < 3 {
		return nil
	}
	var e RawEvent
	if err := json.Unmarshal((*m.RawMessage)[2], &e); err != nil {
		return nil
	}
	return &e
}

func (m *EventMessage) SubscriptionID() []byte {
	if len(*m.RawMessage) < 3 {
		return []byte{}
	}
	var sid string
	if err := json.Unmarshal((*m.RawMessage)[1], &sid); err != nil {
		return []byte{}
	}
	return []byte(sid)
}

func NewNoticeMessage() Message {
	mess := &NoticeMessage{&RawMessage{}}
	mess.Push(MessageTypeNotice)
	return mess
}

// NoticeMessage TBD
type NoticeMessage struct {
	*RawMessage
}

func NewOkMessage() Message {
	mess := &OkMessage{}
	mess.Push(MessageTypeOk)
	return mess
}

// OkMessage TBD
type OkMessage struct {
	*RawMessage
}

func NewRequestMessage(sid string, filters ...*Filter) Message {
	mess := &RequestMessage{&RawMessage{}}
	mess.Push(MessageTypeRequest)
	mess.Push(sid)
	for _, f := range filters {
		mess.Push(f)
	}
	return mess
}

// RequestMessage TBD
type RequestMessage struct {
	*RawMessage
}

// SubscriptionID TBD
func (m *RequestMessage) SubscriptionID() []byte {
	if len(*m.RawMessage) < 2 {
		return []byte{}
	}
	var sid string
	if err := json.Unmarshal((*m.RawMessage)[1], &sid); err != nil {
		return []byte{}
	}
	return []byte(sid)
}
