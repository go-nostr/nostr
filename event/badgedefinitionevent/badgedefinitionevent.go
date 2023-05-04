package badgedefinitionevent

import (
	"github.com/go-nostr/nostr/event"
	"github.com/go-nostr/nostr/tag"
)

const Kind = 30009

func New(content string, tags ...tag.Tag) *event.Event {
	return event.New(Kind, content, tags...)
}
