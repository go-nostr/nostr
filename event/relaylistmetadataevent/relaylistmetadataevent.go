package relaylistmetadataevent

import "github.com/go-nostr/nostr/event"

// Kind for managing relay list metadata
const Kind = 10002

// NewRelayListMetadataEvent creates a new RelayListMetadataEvent.
func NewRelayListMetadataEvent() *RelayListMetadataEvent {
	event := &RelayListMetadataEvent{}
	event.Set("kind", Kind)
	return event
}

// RelayListMetadata represents a relay list metadata event.
type RelayListMetadataEvent struct {
	*event.Event
}
