package pubkeytag

import "github.com/go-nostr/nostr/tag"

const Type = "p"

func New(pubkey string, relay string, petname string) *tag.Tag {
	return tag.New(Type, pubkey, relay, petname)
}
