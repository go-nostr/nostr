package metadataevent

import (
	"github.com/go-nostr/nostr/event"
)

const Kind = 0 // Event kind for setting metadata

// New creates a new metadata event.
func New() *MetadataEvent {
	evnt := &MetadataEvent{&event.Event{}}
	evnt.Set("kind", Kind)
	return evnt
}

// MetadataEvent represents a metadata event.
type MetadataEvent struct {
	*event.Event
}
