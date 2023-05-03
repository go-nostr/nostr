package nostr

import (
	"encoding/json"
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
	Push(val any) error
	Type() []byte
	Unmarshal(data []byte) error
	Values() []any
}
