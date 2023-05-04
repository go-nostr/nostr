package challengetag

import "github.com/go-nostr/nostr/tag"

const Type = "challenge"

func New(challenge string) tag.Tag {
	return tag.Tag{
		Type,
		challenge,
	}
}
