package countmessage

import (
	"github.com/go-nostr/nostr/message"
)

type Count struct {
	Count int `json:"count,omitempty"`
}

const Type = "COUNT"

type Options struct {
	Count          int
	SubscriptionID string
}

// New creates a new CountMessage.
func New(opt *Options) (*message.Message, error) {
	if opt.SubscriptionID != "" {
		return message.New(Type, opt.SubscriptionID), nil
	}
	return message.New(Type, &Count{
		Count: opt.Count,
	}), nil
}
