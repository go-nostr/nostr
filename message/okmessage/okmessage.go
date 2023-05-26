package okmessage

import "github.com/go-nostr/nostr/message"

const Type = "OK"

// New creates a new OkMessage.
func New(eventID string, ok bool, status string) message.Message {
	return message.New(Type, eventID, ok, status)
}
