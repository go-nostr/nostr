package shorttextnote

import (
	"github.com/go-nostr/nostr/event"
)

const Kind = 1

func New(content string) *event.Event {
	return event.New(Kind, content)
}
