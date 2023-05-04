package badgeawardevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for awarding badges to users
const Kind = 8

// New creates a new badge award event.
func New(content string) *event.Event {
	return event.New(Kind, content)
}
