package contactsevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for contact List and petnames
const Kind = 3

// New creates a new contacts event.
func New() *event.Event {
	return event.New(Kind, "")
}
