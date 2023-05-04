package channelhidemessageevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for hiding messages in a channel
const Kind = 43

// New TODO
func New() *event.Event {
	return event.New(Kind, "")
}
