package badgenametag

import "github.com/go-nostr/nostr/tag"

const Type = "name"

func New(name string) tag.Tag {
	return tag.Tag{
		Type,
		name,
	}
}
