package categorizedbookmarklistevent

import (
	"github.com/go-nostr/nostr/event"
)

// Kind for managing categorized bookmark list
const Kind = 30001

type Options struct{}

// New creates a new CategorizedBookmarkListEvent.
func New(opt *Options) *CategorizedBookmarkListEvent {
	return &CategorizedBookmarkListEvent{
		Event: event.New(&event.Options{
			Kind: Kind,
		}),
	}
}

// CategorizedBookmarkList represents a categorized bookmark list event.
type CategorizedBookmarkListEvent struct {
	*event.Event
}
