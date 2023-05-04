package tag

import (
	"encoding/json"
)

func New(v ...any) Tag {
	if v == nil {
		v = make([]any, 0)
	}
	return Tag(v)
}

type Tag []any

func (t *Tag) Push(v any) {
	if t == nil {
		*t = make(Tag, 0)
	}
	*t = append(*t, v)
}

func (t Tag) Marshal() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Tag) Unmarshal(data []byte) error {
	return json.Unmarshal(data, t)
}

func (t Tag) Values() []any {
	return t
}
