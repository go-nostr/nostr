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
func New(name string, about string, picture string) (*event.Event, error) {
	metadata := &Metadata{
		About:   about,
		Name:    name,
		Picture: picture,
	}
	data, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}
	return event.New(Kind, string(data)), nil
}

func GetMetadata(evnt *event.Event) (*Metadata, error) {
	var metadata Metadata
	if err := json.Unmarshal([]byte(evnt.Content), &metadata); err != nil {
		return nil, err
	}
	return &metadata, nil
}

func Validate(evnt *event.Event) error {
	if _, err := GetMetadata(evnt); err != nil {
		return err
	}
	if err := evnt.Verify(); err != nil {
		return err
	}
	return nil
}
