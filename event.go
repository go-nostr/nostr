package nostr

// EventKind represents the different types of events available.
type EventKind int64

const (
	EventKindMetadata                EventKind = 0     // Event for setting metadata
	EventKindShortTextNote           EventKind = 1     // Event for storing a text note
	EventKindRecommendRelay          EventKind = 2     // Event for recommending a server
	EventKindContacts                EventKind = 3     // Event for contact List and petnames
	EventKindEncryptedDirectMessages EventKind = 4     // Event for encrypted direct messages
	EventKindEventDeletion           EventKind = 5     // Event for deleting events
	EventKindReposts                 EventKind = 6     // Event for signaling to followers that another event is worth reading
	EventKindReaction                EventKind = 7     // Event for reacting to other notes
	EventKindBadgeAward              EventKind = 8     // Event for awarding badges to users
	EventKindChannelCreation         EventKind = 40    // Event for creating new channels
	EventKindChannelMetadata         EventKind = 41    // Event for setting channel metadata
	EventKindChannelMessage          EventKind = 42    // Event for posting messages in a channel
	EventKindChannelHideMessage      EventKind = 43    // Event for hiding messages in a channel
	EventKindChannelMuteUser         EventKind = 44    // Event for muting a user in a channel
	EventKindReporting               EventKind = 1984  // Event for reporting content or users
	EventKindZapRequest              EventKind = 9734  // Event for requesting a Zap action
	EventKindZap                     EventKind = 9735  // Event for performing a Zap action
	EventKindMuteList                EventKind = 10000 // Event for managing a mute list
	EventKindPinList                 EventKind = 10001 // Event for managing a pin list
	EventKindRelayListMetadata       EventKind = 10002 // Event for managing relay list metadata
	EventKindClientAuthentication    EventKind = 22242 // Event for client authentication process
	EventKindNostrConnect            EventKind = 24133 // Event for Nostr client connection
	EventKindCategorizedPeopleList   EventKind = 30000 // Event for managing categorized people list
	EventKindCategorizedBookmarkList EventKind = 30001 // Event for managing categorized bookmark list
	EventKindProfileBadges           EventKind = 30008 // Event for profile badges management
	EventKindBadgeDefinition         EventKind = 30009 // Event for defining badges
	EventKindCreateOrUpdateStall     EventKind = 30017 // Event for creating or updating a stall
	EventKindCreateOrUpdateProduct   EventKind = 30018 // Event for creating or updating a product
	EventKindLongFormContent         EventKind = 30023 // Event for posting long-form content
	EventKindApplicationSpecificData EventKind = 30078 // Event for managing application-specific data
)

// Event represents an event with a specified kind.
type Event interface {
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
}
