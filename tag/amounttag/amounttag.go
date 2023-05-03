package amounttag

import "github.com/go-nostr/nostr/tag"

const Type = "amount"

func New() *AmountTag {
	tag := &AmountTag{&tag.Tag{}}
	tag.Push(Type)
	return tag
}

type AmountTag struct {
	*tag.Tag
}
