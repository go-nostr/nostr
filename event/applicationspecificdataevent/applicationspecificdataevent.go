package applicationspecificdataevent

import (
	"github.com/go-nostr/nostr/event"
)

const Kind = 30078

func New(content string) *event.Event {
	return event.New(Kind, content)
}
