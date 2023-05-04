package eventtag

import "github.com/go-nostr/nostr/tag"

const Type = "a"

func New(naddr string) *tag.Tag {
	return tag.New(Type, naddr)
}
