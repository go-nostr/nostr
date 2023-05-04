package identifiertag

import "github.com/go-nostr/nostr/tag"

const Type = "d"

func New(identifier string) *tag.Tag {
	return tag.New(Type, identifier)
}
