package okmessage

import "github.com/go-nostr/nostr/message"

const Type = "OK"

// New creates a new OkMessage.
func New(id string, ok bool, msg string) message.Message {
	return message.New(Type, id, ok, msg)
}
