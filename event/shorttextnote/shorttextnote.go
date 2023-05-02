package shorttextnote

import (
	"github.com/go-nostr/nostr/event"
)

const Kind = 1

// New creates a new short text note event.
func New(content string) *ShortTextNoteEvent {
	event := &ShortTextNoteEvent{&event.Event{}}
	event.Set("kind", Kind)
	event.Set("content", content)
	return event
}

// ShortTextNoteEvent represents a short text note event.
type ShortTextNoteEvent struct {
	*event.Event
}
