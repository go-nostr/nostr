package rawtag

import "encoding/json"

// New creates a new RawTag with the given type.
func New(typ string) *RawTag {
	tag := &RawTag{}
	tag.Push(typ)
	return tag
}

// RawTag is a raw representation of a Tag as a slice of json.RawMessage.
type RawTag []json.RawMessage

// Marshal marshals the RawTag into a JSON byte slice.
func (t *RawTag) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

// Push appends a value to the RawTag after marshaling it into a JSON RawMessage.
func (t *RawTag) Push(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	*t = append(*t, data)
	return nil
}

// Type returns the type of the RawTag.
func (t *RawTag) Type() string {
	if len(*t) < 1 {
		return ""
	}
	typ, ok := t.Values()[0].(string)
	if !ok {
		return ""
	}
	return typ
}

// Unmarshal unmarshals a JSON byte slice into a RawTag.
func (t *RawTag) Unmarshal(data []byte) error {
	return json.Unmarshal(data, t)
}

// Values returns the values of the RawTag as a slice of any.
func (t *RawTag) Values() []any {
	var vals []any
	for _, arg := range *t {
		var val any
		json.Unmarshal(arg, &val)
		vals = append(vals, val)
	}
	return vals
}
