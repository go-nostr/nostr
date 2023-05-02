package message

import (
	"encoding/json"
	"fmt"
)

// New creates a new Message with the provided type.
func New(typ string) *Message {
	mess := &Message{}
	mess.Push(typ)
	return mess
}

// Message is a raw representation of a Message as a slice of json.Message.
type Message []json.RawMessage

// Marshal marshals the Message into a JSON byte slice.
func (m *Message) Marshal() ([]byte, error) {
	return m.MarshalJSON()
}

// MarshalJSON TBD
func (m *Message) MarshalJSON() ([]byte, error) {
	byt := make([]byte, 0)
	byt = append(byt, '[')
	len := len(*m)
	for i, v := range *m {
		data, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		byt = append(byt, data...)
		if i < len-1 {
			byt = append(byt, ',')
		}
	}
	byt = append(byt, ']')
	return byt, nil
}

// Push appends a value to the Message after marshaling it into a JSON Message.
func (m *Message) Push(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("%s", err)
		return err
	}
	*m = append(*m, data)
	return nil
}

// Type returns the type of the Message.
func (m *Message) Type() []byte {
	if len(*m) < 1 {
		return []byte{}
	}
	var typ string
	if err := json.Unmarshal((*m)[0], &typ); err != nil {
		return []byte{}
	}
	return []byte(typ)
}

// Unmarshal unmarshals a JSON byte slice into a Message.
func (m *Message) Unmarshal(data []byte) error {
	return m.UnmarshalJSON(data)
}

// UnmarshalJSON TBD
func (m *Message) UnmarshalJSON(data []byte) error {
	mess := make([]json.RawMessage, 0)
	if err := json.Unmarshal(data, &mess); err != nil {
		return err
	}
	*m = mess
	return nil
}

// Values returns the values of the Message as a slice of any.
func (m *Message) Values() []any {
	var vals []any
	for _, arg := range *m {
		var val any
		json.Unmarshal(arg, &val)
		vals = append(vals, val)
	}
	return vals
}
