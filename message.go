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
	json.Marshaler
	json.Unmarshaler
	Marshal() ([]byte, error)
	Push(v any) error
	Type() []byte
	Unmarshal(data []byte) error
	Values() []any
}
