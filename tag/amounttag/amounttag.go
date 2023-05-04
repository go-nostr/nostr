package amounttag

import (
	"github.com/go-nostr/nostr/tag"
)

const Type = "amount"

func New(amount int) *tag.Tag {
	return &tag.Tag{
		Type,
		amount,
	}
}
