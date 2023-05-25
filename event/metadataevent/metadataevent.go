package metadataevent

import (
	"encoding/json"

	"github.com/go-nostr/nostr/event"
)

type Metadata struct {
	About   string `json:"about,omitempty"`
	Name    string `json:"name,omitempty"`
	Picture string `json:"picture,omitempty"`
}

// Kind for metadata
const Kind = 0

// New creates a new metadata event
func New(name string, about string, picture string) *event.Event {
	metadata := &Metadata{
		About:   about,
		Name:    name,
		Picture: picture,
	}
	data, err := json.Marshal(metadata)
	if err != nil {
		panic(err)
	}
	return event.New(Kind, string(data))
}

func GetMetadata(evt *event.Event) (*Metadata, error) {
	var metadata Metadata
	if err := json.Unmarshal([]byte(evt.Content), &metadata); err != nil {
		return nil, err
	}
	return &metadata, nil
}

func Validate(evt *event.Event) error {
	if _, err := GetMetadata(evt); err != nil {
		return err
	}
	if err := evt.Verify(); err != nil {
		return err
	}
	return nil
}
