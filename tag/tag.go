package tag

import "encoding/json"

// New creates a new Tag with the given type.
func New(typ string) *Tag {
	tag := &Tag{}
	tag.Push(typ)
	return tag
}

// Tag is a raw representation of a Tag as a slice of json.RawMessage.
type Tag []json.RawMessage

// Marshal marshals the Tag into a JSON byte slice.
func (t *Tag) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

// Push appends a value to the Tag after marshaling it into a JSON RawMessage.
func (t *Tag) Push(v any) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	*t = append(*t, data)
	return nil
}

// Type returns the type of the Tag.
func (t *Tag) Type() string {
	if len(*t) < 1 {
		return ""
	}
	typ, ok := t.Values()[0].(string)
	if !ok {
		return ""
	}
	return typ
}

// Unmarshal unmarshals a JSON byte slice into a Tag.
func (t *Tag) Unmarshal(data []byte) error {
	return json.Unmarshal(data, t)
}

// Values returns the values of the Tag as a slice of any.
func (t *Tag) Values() []any {
	var vals []any
	for _, arg := range *t {
		var val any
		json.Unmarshal(arg, &val)
		vals = append(vals, val)
	}
	return vals
}
