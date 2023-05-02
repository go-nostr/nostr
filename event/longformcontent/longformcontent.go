package longformcontent

import "github.com/go-nostr/nostr/event"

// Kind for posting long-form content
const Kind = 30023

// New creates a new LongFormContentEvent.
func New() *LongFormContentEvent {
	event := &LongFormContentEvent{&event.Event{}}
	event.Set("kind", Kind)
	return event
}

// LongFormContent represents a long form content event.
type LongFormContentEvent struct {
	*event.Event
}
