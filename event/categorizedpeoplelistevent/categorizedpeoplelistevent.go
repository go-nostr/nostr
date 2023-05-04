package categorizedpeoplelistevent

import (
	"github.com/go-nostr/nostr/event"
	"github.com/go-nostr/nostr/tag"
)

// Kind for managing categorized people list
const Kind = 30000

// New creates a new CategorizedPeopleListEvent.
func New(content string, tags ...tag.Tag) *event.Event {
	return event.New(Kind, content, tags...)
}
