package authmessage

import (
	"github.com/go-nostr/nostr/event"
	"github.com/go-nostr/nostr/message"
)

// Type for message "AUTH"
const Type = "AUTH"

// Options for creating a new "AUTH" message
type Options struct {
	Challenge string
	Event     *event.Event
}

// New creates a new "AUTH" message
func New(opt *Options) message.Message {
	if opt.Challenge != "" {
		return message.New(Type, opt.Challenge)
	}
	return message.New(Type, opt.Event)
}
