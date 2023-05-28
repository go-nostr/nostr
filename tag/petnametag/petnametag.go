package petnametag

import "github.com/go-nostr/nostr/tag"

const Type = "p"

func New(pubKeyStr string, relayURL string, petname string) tag.Tag {
	return tag.New(Type, pubKeyStr, relayURL, petname)
}
