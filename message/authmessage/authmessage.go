package authmessage

import (
	"github.com/go-nostr/nostr/event"
	"github.com/go-nostr/nostr/message"
)

// Type TODO
const Type = "AUTH"

// New creates a new AuthMessage with the provided challenge and event.
func New(opt *Options) *AuthMessage {
	mess := &AuthMessage{
		Message: &message.Message{},
	}
	mess.Push(Type)
	if opt.Challenge != "" {
		mess.Push(opt.Challenge)
	}
	if opt.Event != nil {
		mess.Push(opt.Event)
	}
	return mess
}

type Options struct {
	Challenge string
	Event     event.Event
}

// AuthMessage is a specialized message type for authentication.
type AuthMessage struct {
	*message.Message
}
