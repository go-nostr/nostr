package noticemessage

import "github.com/go-nostr/nostr/message"

const Type = "NOTICE"

// New creates a new NoticeMessage.
func New() *NoticeMessage {
	mess := &NoticeMessage{&message.Message{}}
	mess.Push(Type)
	return mess
}

// NoticeMessage is a specialized message type for sending notifications.
type NoticeMessage struct {
	*message.Message
}
