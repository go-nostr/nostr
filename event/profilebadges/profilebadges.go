package profilebadges

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for profile badges management
const Kind = 30008

// New creates a new ProfileBadgesEvent.
func New() *ProfileBadgesEvent {
	event := &ProfileBadgesEvent{}
	event.Set("kind", Kind)
	return event
}

// ProfileBadges represents a profile badges event.
type ProfileBadgesEvent struct {
	*event.Event
}
