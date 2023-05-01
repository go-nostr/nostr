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

// MessageHandlerFunc is a function type that takes a Message as a parameter.
type MessageHandlerFunc func(mess Message)

// Handle calls the MessageHandlerFunc with the provided Message.
func (f MessageHandlerFunc) Handle(mess Message) {
	f(mess)
}

// MessageHandler is an interface for handling Message types.
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

// NewRawMessage creates a new RawMessage with the provided type.
func NewRawMessage(typ string) Message {
	mess := &RawMessage{}
	mess.Push(typ)
	return mess
}

// RawMessage is a raw representation of a Message as a slice of json.RawMessage.
type RawMessage []json.RawMessage

// Marshal marshals the RawMessage into a JSON byte slice.
func (m *RawMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Push appends a value to the RawMessage after marshaling it into a JSON RawMessage.
func (m *RawMessage) Push(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	*m = append(*m, data)
	return nil
}

// Type returns the type of the RawMessage.
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

// Unmarshal unmarshals a JSON byte slice into a RawMessage.
func (m *RawMessage) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Values returns the values of the RawMessage as a slice of any.
func (m *RawMessage) Values() []any {
	var vals []any
	for _, arg := range *m {
		var val any
		json.Unmarshal(arg, &val)
		vals = append(vals, val)
	}
	return vals
}

// NewAuthMessage creates a new AuthMessage with the provided challenge and event.
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

// AuthMessage is a specialized message type for authentication.
type AuthMessage struct {
	*RawMessage
}

// NewCloseMessage creates a new CloseMessage.
func NewCloseMessage() Message {
	mess := &CloseMessage{&RawMessage{}}
	mess.Push(MessageTypeClose)
	return mess
}

// CloseMessage is a specialized message type for closing a connection.
type CloseMessage struct {
	*RawMessage
}

// NewCountMessage creates a new CountMessage.
func NewCountMessage() Message {
	mess := &CountMessage{&RawMessage{}}
	mess.Push(MessageTypeCount)
	return mess
}

// CountMessage is a specialized message type for counting events.
type CountMessage struct {
	*RawMessage
}

// NewEOSEMessage creates a new EOSEMessage.
func NewEOSEMessage() Message {
	mess := &EOSEMessage{&RawMessage{}}
	mess.Push(MessageTypeEOSE)
	return mess
}

// EOSEMessage is a specialized message type for indicating the end of stored events.
type EOSEMessage struct {
	*RawMessage
}

// SubscriptionID returns the subscription ID contained in the RequestMessage.
func (m *EOSEMessage) SubscriptionID() []byte {
	if len(*m.RawMessage) < 2 {
		return []byte{}
	}
	var sid string
	if err := json.Unmarshal((*m.RawMessage)[1], &sid); err != nil {
		return []byte{}
	}
	return []byte(sid)
}

// NewEventMessage creates a new EventMessage.
func NewEventMessage() Message {
	mess := &EventMessage{&RawMessage{}}
	mess.Push(MessageTypeEvent)
	return mess
}

// EventMessage is a specialized message type for handling events.
type EventMessage struct {
	*RawMessage
}

// Event returns the event contained in the EventMessage.
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

// SubscriptionID returns the subscription ID contained in the EventMessage.
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

// NewNoticeMessage creates a new NoticeMessage.
func NewNoticeMessage() Message {
	mess := &NoticeMessage{&RawMessage{}}
	mess.Push(MessageTypeNotice)
	return mess
}

// NoticeMessage is a specialized message type for sending notifications.
type NoticeMessage struct {
	*RawMessage
}

// NewOkMessage creates a new OkMessage.
func NewOkMessage() Message {
	mess := &OkMessage{}
	mess.Push(MessageTypeOk)
	return mess
}

// OkMessage is a specialized message type for indicating success or confirmation.
type OkMessage struct {
	*RawMessage
}

// NewRequestMessage creates a new RequestMessage with the provided subscription ID and filters.
func NewRequestMessage(sid string, filters ...*Filter) Message {
	mess := &RequestMessage{&RawMessage{}}
	mess.Push(MessageTypeRequest)
	mess.Push(sid)
	for _, f := range filters {
		mess.Push(f)
	}
	return mess
}

// RequestMessage is a specialized message type for making requests.
type RequestMessage struct {
	*RawMessage
}

// SubscriptionID returns the subscription ID contained in the RequestMessage.
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
