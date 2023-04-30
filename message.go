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
	Marshal() ([]byte, error)
	Type() []byte
	Unmarshal(data []byte) error
	Values() []any
}

func NewRawMessage(typ string, val ...any) Message {
	mess := RawMessage{}
	mess.Push(typ)
	for _, v := range val {
		mess.Push(v)
	}
	return mess
}

// RawMessage TBD
type RawMessage []json.RawMessage

// Marshal TBD
func (m RawMessage) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Args TBD
func (m *RawMessage) Push(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	*m = append(*m, data)
	return nil
}

// Type TBD
func (m RawMessage) Type() []byte {
	var t string
	if err := json.Unmarshal(m[0], &t); err != nil {
		return nil
	}
	return []byte(t)
}

// Unmarshal TBD
func (m RawMessage) Unmarshal(data []byte) error {
	return json.Unmarshal(data, &m)
}

// Values TBD
func (m RawMessage) Values() []any {
	v := make([]any, len(m))
	for i := 0; i < len(m); i++ {
		err := json.Unmarshal(m[i], v[i])
		if err != nil {
			continue
		}
	}
	return v
}

// AuthMessage TBD
type AuthMessage = RawMessage

// NewAuthMessage TBD
func NewAuthMessage(challenge string, event *Event) Message {
	mess := AuthMessage{}
	mess.Push(MessageTypeAuth)
	if challenge != "" {
		mess.Push(challenge)
	} else if event != nil {
		mess.Push(event)

	}
	return mess
}

// CloseMessage TBD
type CloseMessage = RawMessage

// NewCloseMessage TBD
func NewCloseMessage(subscriptionID string) Message {
	mess := CloseMessage{}
	mess.Push(MessageTypeClose)
	mess.Push(subscriptionID)
	return mess
}

// CountMessage TBD
type CountMessage = RawMessage

// NewCountMessage TBD
func NewCountMessage(subscriptionID string, count *Count, filters ...*Filter) Message {
	mess := CountMessage{}
	mess.Push(MessageTypeCount)
	mess.Push(subscriptionID)
	if count != nil {
		mess.Push(count)
		return mess
	}
	for _, v := range filters {
		mess.Push(v)
	}
	return mess
}

// EOSEMessage TBD
type EOSEMessage = RawMessage

// NewEOSEMessage TBD
func NewEOSEMessage(subscriptionID string) Message {
	mess := EOSEMessage{}
	mess.Push(MessageTypeEOSE)
	mess.Push(subscriptionID)
	return mess
}

// EventMessage TBD
type EventMessage = RawMessage

// NewEventMessage TBD
func NewEventMessage(subscriptionID string, event Event) Message {
	mess := EventMessage{}
	mess.Push(MessageTypeEvent)
	mess.Push(subscriptionID)
	mess.Push(event)
	return mess
}

func (m EventMessage) Event() Event {
	var e RawEvent
	if err := json.Unmarshal(m[2], &e); err != nil {
		return nil
	}
	return &e
}

func (m EventMessage) SubscriptionID() []byte {
	var subscriptionID []byte
	if err := json.Unmarshal(m[1], &subscriptionID); err != nil {
		return nil
	}
	return subscriptionID
}

// NoticeMessage TBD
type NoticeMessage = RawMessage

// NewNoticeMessage TBD
func NewNoticeMessage(notice string) Message {
	mess := NoticeMessage{}
	mess.Push(MessageTypeNotice)
	mess.Push(notice)
	return mess
}

// OkMessage TBD
type OkMessage = RawMessage

// NewOkMessage TBD
func NewOkMessage(eventID string, isSaved bool, message string) Message {
	mess := OkMessage{}
	mess.Push(MessageTypeOk)
	mess.Push(eventID)
	mess.Push(isSaved)
	mess.Push(message)
	return mess
}

// RequestMessage TBD
type RequestMessage = RawMessage

// NewRequestMessage TBD
func NewRequestMessage(subscriptionID string, filter *Filter) Message {
	mess := RequestMessage{}
	mess.Push(MessageTypeRequest)
	mess.Push(subscriptionID)
	mess.Push(filter)
	return mess
}
