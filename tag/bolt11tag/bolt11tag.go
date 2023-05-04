package bolt11tag

import "github.com/go-nostr/nostr/tag"

const Type = "bolt11"

// New tag containing the description hash bolt11 invoice.
func New(hash string) tag.Tag {
	return tag.Tag{
		Type,
		hash,
	}
}
