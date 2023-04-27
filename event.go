package nostr

import "encoding/json"

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

// newBaseEvent TBD
func newBaseEvent() Event {
	return &BaseEvent{}
}

type BaseEvent struct {
	ID        string    `json:"id,omitempty"`
	PubKey    string    `json:"pub_key,omitempty"`
	CreatedAt int64     `json:"created_at,omitempty"`
	Kind      EventKind `json:"kind,omitempty"`
	Tags      []Tag     `json:"tags,omitempty"`
	Content   string    `json:"content,omitempty"`
	Signature string    `json:"signature,omitempty"`
}

// Marshal TBD
func (e *BaseEvent) Marshal() ([]byte, error) {
	return json.Marshal(e)
}

// Unmarshal TBD
func (e *BaseEvent) Unmarshal(data []byte) error {
	return json.Unmarshal(data, e)
}

// Validate TBD
func (e *BaseEvent) Validate() error {
	return nil
}

// { TBD
// NewMetadataEvent TBD
func NewMetadataEvent() Event {
	return newBaseEvent()
}

// MetadataEvent TBD
type MetadataEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *MetadataEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *MetadataEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *MetadataEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewShortTextNoteEvent TBD
func NewShortTextNoteEvent() Event {
	return newBaseEvent()
}

// ShortTextNoteEvent TBD
type ShortTextNoteEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ShortTextNoteEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ShortTextNoteEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ShortTextNoteEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewRecommendRelayEvent TBD
func NewRecommendRelayEvent() Event {
	return newBaseEvent()
}

// RecommendRelayEvent TBD
type RecommendRelayEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *RecommendRelayEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *RecommendRelayEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *RecommendRelayEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewContactsEvent TBD
func NewContactsEvent() Event {
	return newBaseEvent()
}

// ContactsEvent TBD
type ContactsEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ContactsEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ContactsEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ContactsEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewEncryptedDirectMessagesEvent TBD
func NewEncryptedDirectMessagesEvent() Event {
	return newBaseEvent()
}

// EncryptedDirectMessagesEvent TBD
type EncryptedDirectMessagesEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *EncryptedDirectMessagesEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *EncryptedDirectMessagesEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *EncryptedDirectMessagesEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewEventDeletionEvent TBD
func NewEventDeletionEvent() Event {
	return newBaseEvent()
}

// EventDeletionEvent TBD
type EventDeletionEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *EventDeletionEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *EventDeletionEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *EventDeletionEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewEventRepostsEvent TBD
func NewEventRepostsEvent() Event {
	return newBaseEvent()
}

// EventRepostsEvent TBD
type EventRepostsEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *EventRepostsEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *EventRepostsEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *EventRepostsEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewReactionEvent TBD
func NewReactionEvent() Event {
	return newBaseEvent()
}

// ReactionEvent TBD
type ReactionEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ReactionEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ReactionEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ReactionEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewBadgeAwardEvent TBD
func NewBadgeAwardEvent() Event {
	return newBaseEvent()
}

// BadgeAwardEvent TBD
type BadgeAwardEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *BadgeAwardEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *BadgeAwardEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *BadgeAwardEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// NewChannelCreationEvent TBD
func NewChannelCreationEvent() Event {
	return newBaseEvent()
}

// ChannelCreationEvent TBD
type ChannelCreationEvent struct {
	*BaseEvent
}

// Marshal TBD
func (e *ChannelCreationEvent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ChannelCreationEvent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ChannelCreationEvent) Validate() error {
	return e.BaseEvent.Validate()
}

// ChannelCreation TBD
type ChannelCreation struct {
	*BaseEvent
}

// Marshal TBD
func (e *ChannelCreation) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ChannelCreation) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ChannelCreation) Validate() error {
	return e.BaseEvent.Validate()
}

// ChannelMetadata TBD
type ChannelMetadata struct {
	*BaseEvent
}

// Marshal TBD
func (e *ChannelMetadata) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ChannelMetadata) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ChannelMetadata) Validate() error {
	return e.BaseEvent.Validate()
}

// ChannelMessage TBD
type ChannelMessage struct {
	*BaseEvent
}

// Marshal TBD
func (e *ChannelMessage) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ChannelMessage) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ChannelMessage) Validate() error {
	return e.BaseEvent.Validate()
}

// ChannelHideMessage TBD
type ChannelHideMessage struct {
	*BaseEvent
}

// Marshal TBD
func (e *ChannelHideMessage) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ChannelHideMessage) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ChannelHideMessage) Validate() error {
	return e.BaseEvent.Validate()
}

// ChannelMuteUser TBD
type ChannelMuteUser struct {
	*BaseEvent
}

// Marshal TBD
func (e *ChannelMuteUser) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ChannelMuteUser) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ChannelMuteUser) Validate() error {
	return e.BaseEvent.Validate()
}

// Reporting TBD
type Reporting struct {
	*BaseEvent
}

// Marshal TBD
func (e *Reporting) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *Reporting) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *Reporting) Validate() error {
	return e.BaseEvent.Validate()
}

// ZapRequest TBD
type ZapRequest struct {
	*BaseEvent
}

// Marshal TBD
func (e *ZapRequest) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ZapRequest) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ZapRequest) Validate() error {
	return e.BaseEvent.Validate()
}

// Zap TBD
type Zap struct {
	*BaseEvent
}

// Marshal TBD
func (e *Zap) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *Zap) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *Zap) Validate() error {
	return e.BaseEvent.Validate()
}

// MuteList TBD
type MuteList struct {
	*BaseEvent
}

// Marshal TBD
func (e *MuteList) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *MuteList) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *MuteList) Validate() error {
	return e.BaseEvent.Validate()
}

// PinList TBD
type PinList struct {
	*BaseEvent
}

// Marshal TBD
func (e *PinList) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *PinList) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *PinList) Validate() error {
	return e.BaseEvent.Validate()
}

// RelayListMetadata TBD
type RelayListMetadata struct {
	*BaseEvent
}

// Marshal TBD
func (e *RelayListMetadata) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *RelayListMetadata) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *RelayListMetadata) Validate() error {
	return e.BaseEvent.Validate()
}

// ClientAuthentication TBD
type ClientAuthentication struct {
	*BaseEvent
}

// Marshal TBD
func (e *ClientAuthentication) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ClientAuthentication) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ClientAuthentication) Validate() error {
	return e.BaseEvent.Validate()
}

// NostrConnect TBD
type NostrConnect struct {
	*BaseEvent
}

// Marshal TBD
func (e *NostrConnect) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *NostrConnect) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *NostrConnect) Validate() error {
	return e.BaseEvent.Validate()
}

// CategorizedPeopleList TBD
type CategorizedPeopleList struct {
	*BaseEvent
}

// Marshal TBD
func (e *CategorizedPeopleList) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *CategorizedPeopleList) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *CategorizedPeopleList) Validate() error {
	return e.BaseEvent.Validate()
}

// CategorizedBookmarkList TBD
type CategorizedBookmarkList struct {
	*BaseEvent
}

// Marshal TBD
func (e *CategorizedBookmarkList) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *CategorizedBookmarkList) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *CategorizedBookmarkList) Validate() error {
	return e.BaseEvent.Validate()
}

// ProfileBadges TBD
type ProfileBadges struct {
	*BaseEvent
}

// Marshal TBD
func (e *ProfileBadges) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ProfileBadges) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ProfileBadges) Validate() error {
	return e.BaseEvent.Validate()
}

// BadgeDefinition TBD
type BadgeDefinition struct {
	*BaseEvent
}

// Marshal TBD
func (e *BadgeDefinition) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *BadgeDefinition) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *BadgeDefinition) Validate() error {
	return e.BaseEvent.Validate()
}

// CreateOrUpdateStall TBD
type CreateOrUpdateStall struct {
	*BaseEvent
}

// Marshal TBD
func (e *CreateOrUpdateStall) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *CreateOrUpdateStall) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *CreateOrUpdateStall) Validate() error {
	return e.BaseEvent.Validate()
}

// CreateOrUpdateProduct TBD
type CreateOrUpdateProduct struct {
	*BaseEvent
}

// Marshal TBD
func (e *CreateOrUpdateProduct) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *CreateOrUpdateProduct) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *CreateOrUpdateProduct) Validate() error {
	return e.BaseEvent.Validate()
}

// LongFormContent TBD
type LongFormContent struct {
	*BaseEvent
}

// Marshal TBD
func (e *LongFormContent) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *LongFormContent) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *LongFormContent) Validate() error {
	return e.BaseEvent.Validate()
}

// ApplicationSpecificData TBD
type ApplicationSpecificData struct {
	*BaseEvent
}

// Marshal TBD
func (e *ApplicationSpecificData) Marshal() ([]byte, error) {
	return e.BaseEvent.Marshal()
}

// Unmarshal TBD
func (e *ApplicationSpecificData) Unmarshal(data []byte) error {
	return e.BaseEvent.Unmarshal(data)
}

// Validate TBD
func (e *ApplicationSpecificData) Validate() error {
	return e.BaseEvent.Validate()
}
