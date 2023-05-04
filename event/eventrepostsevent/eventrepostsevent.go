package eventrepostsevent

import (
	"github.com/go-nostr/nostr/event"
)

const Kind = 6

func New() *event.Event {
	return event.New(Kind, "")
}
