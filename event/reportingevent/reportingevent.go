package reportingevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for reporting content or users
const Kind = 1984

// New creates a new ReportingEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
