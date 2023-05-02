package categorizedpeoplelistevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for managing categorized people list
const Kind = 30000

// New creates a new CategorizedPeopleListEvent.
func New() *CategorizedPeopleListEvent {
	event := &CategorizedPeopleListEvent{&event.Event{}}
	event.Set("kind", Kind)
	return event
}

// CategorizedPeopleList represents a categorized people list event.
type CategorizedPeopleListEvent struct {
	*event.Event
}
