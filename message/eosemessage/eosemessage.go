package eosemessage

import (
	"github.com/go-nostr/nostr/message"
)

const Type = "EOSE"

func New() *message.Message {
	return message.New(Type)
}
