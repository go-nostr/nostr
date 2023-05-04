package message

import (
	"encoding/json"
)

// Handler is an interface for handling Message types.
type Handler interface {
	Handle(mess Message)
}

// HandlerFunc is a function type that takes a Message as a parameter.
type HandlerFunc func(mess Message)

// Handle calls the HandlerFunc with the provided Message.
func (f HandlerFunc) Handle(mess Message) {
	f(mess)
}

// New creates a new Message.
func New(v ...any) Message {
	if v == nil {
		v = make([]any, 0)
	}
	return Message(v)
}

// Message is a raw representation of a Message as a slice of json.Message.
type Message []any

// Marshal marshals the Message into a JSON byte slice.
func (m Message) Marshal() ([]byte, error) {
	return json.Marshal(m)
}

// Push appends a value to the Message after marshaling it into a JSON Message.
func (m *Message) Push(v any) error {
	*m = append(*m, v)
	return nil
}

// Unmarshal unmarshals a JSON byte slice into a Message.
func (m *Message) Unmarshal(data []byte) error {
	return json.Unmarshal(data, m)
}

// Values returns the values of the Message as a slice of any.
func (m Message) Values() []any {
	return []any(m)
}
