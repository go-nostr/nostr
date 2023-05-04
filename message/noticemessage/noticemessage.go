package noticemessage

import "github.com/go-nostr/nostr/message"

const Type = "NOTICE"

// New creates a new NoticeMessage.
func New() *message.Message {
	return message.New(Type)
}
