package profilebadges

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for profile badges management
const Kind = 30008

// New creates a new ProfileBadgesEvent.
func New() *event.Event {
	return event.New(Kind, "")
}
