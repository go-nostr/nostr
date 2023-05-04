package tag

import (
	"encoding/json"
)

func New(v ...any) *Tag {
	t := make(Tag, 0)
	for _, v := range v {
		t = append(t, v)
	}
	return &t
}

type Tag []any

func (t *Tag) Push(v any) {
	if t == nil {
		*t = make(Tag, 0)
	}
	*t = append(*t, v)
}

func (t *Tag) Marshal() ([]byte, error) {
	return json.Marshal(*t)
}

func (t *Tag) Unmarshal(data []byte) error {
	return json.Unmarshal(data, t)
}

func (t *Tag) Values() []any {
	if t == nil {
		*t = make(Tag, 0)
	}
	return *t
}
