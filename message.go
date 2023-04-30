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
func (fn MessageHandlerFunc) Handle(mess Message) {
	fn(mess)
}

// MessageHandler TBD
type MessageHandler interface {
	Handle(mess Message)
}

// Message is an interface for encoding and marshaling messages.
type Message interface {
	Arg(i int) any
	Args() []any
	Marshal() ([]byte, error)
	Type() *string
	Unmarshal(data []byte) error
}

// RawMessage TBD
type RawMessage []any

// Marshal TBD
func (m RawMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Arg TBD
func (m RawMessage) Arg(i int) any {
	if len(m) > i {
		return m[i]
	}
	return nil
}

// Args TBD
func (m RawMessage) Args() []any {
	return m
}

// Type TBD
func (m RawMessage) Type() *string {
	if typ, ok := m.Arg(0).(string); ok {
		return &typ
	}
	return nil
}

// Unmarshal TBD
func (m RawMessage) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &m)
}

// AuthMessage TBD
type AuthMessage = RawMessage

// NewAuthMessage TBD
func NewAuthMessage(challenge string, event *Event) Message {
	if challenge != "" {
		return &AuthMessage{MessageTypeAuth, challenge}
	}
	return &AuthMessage{MessageTypeAuth, event}
}

// CloseMessage TBD
type CloseMessage = RawMessage

// NewCloseMessage TBD
func NewCloseMessage(subscriptionID string) Message {
	return &CloseMessage{MessageTypeClose, subscriptionID}
}

// CountMessage TBD
type CountMessage = RawMessage

// NewCountMessage TBD
func NewCountMessage(subscriptionID string, count *Count, filters ...*Filter) Message {
	if count != nil {
		return &CountMessage{MessageTypeCount, subscriptionID, count}
	}
	mess := &CountMessage{MessageTypeCount, subscriptionID}
	for _, f := range filters {
		*mess = append(*mess, *f)
	}
	return mess
}

// EOSEMessage TBD
type EOSEMessage = RawMessage

// NewEOSEMessage TBD
func NewEOSEMessage(subscriptionID string) Message {
	return &EOSEMessage{
		MessageTypeEOSE,
		subscriptionID,
	}
}

// EventMessage TBD
type EventMessage = RawMessage

// NewEventMessage TBD
func NewEventMessage(subscriptionID string, event Event) Message {
	return &EventMessage{
		MessageTypeEvent,
		subscriptionID,
		event,
	}
}

func (m EventMessage) Event() Event {
	byt, err := json.Marshal(m.Arg(2))
	if err != nil {
		return nil
	}
	var rawEvent RawEvent
	if err := rawEvent.Unmarshal(byt); err != nil {
		return nil
	}
	return &rawEvent
}

// NoticeMessage TBD
type NoticeMessage = RawMessage

// NewNoticeMessage TBD
func NewNoticeMessage(notice string) Message {
	return &NoticeMessage{MessageTypeNotice, notice}
}

// OkMessage TBD
type OkMessage = RawMessage

// NewOkMessage TBD
func NewOkMessage(eventID string, isSaved bool, message string) Message {
	return &OkMessage{MessageTypeOk, eventID, isSaved, message}
}

// RequestMessage TBD
type RequestMessage = RawMessage

// NewRequestMessage TBD
func NewRequestMessage(subscriptionID string, filter *Filter) Message {
	return &RequestMessage{MessageTypeRequest, subscriptionID, filter}
}
