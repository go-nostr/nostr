package recommendrelayevent

import (
	"github.com/go-nostr/nostr/event"
)

const Kind = 2

func New() *event.Event {
	return event.New(Kind, "")
}
