package categorizedbookmarklistevent

import (
	"github.com/go-nostr/nostr/event"
	"github.com/go-nostr/nostr/tag"
)

// Kind for managing categorized bookmark list
const Kind = 30001

// New creates a new CategorizedBookmarkListEvent.
func New(content string, tags ...tag.Tag) *event.Event {
	return event.New(Kind, content, tags...)
}
