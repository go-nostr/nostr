package contactsevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for contact List and petnames
const Kind = 3

// New creates a new contacts event.
func New() *ContactsEvent {
	event := &ContactsEvent{}
	event.Set("kind", Kind)
	return event
}

// ContactsEvent represents a contacts event.
type ContactsEvent struct {
	*event.Event
}
