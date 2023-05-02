package nip

const (
	AuthenticationOfClientsToRelays int = 42 // AuthenticationOfClientsToRelays represents NIP-42: Authentication of clients to relays
	ApplicationSpecificData         int = 78 // ApplicationSpecificData represents NIP-78: Application-specific data
	BasicProtocolFlow               int = 1  // BasicProtocolFlow represents NIP-01: Basic protocol flow description
	BasicKeyDerivation              int = 6  // BasicKeyDerivation represents NIP-06: Basic key derivation from mnemonic seed phrase
	Bech32EncodedEntities           int = 19 // Bech32EncodedEntities represents NIP-19: bech32-encoded entities
	Badges                          int = 58 // Badges represents NIP-58: Badges
	ContactListAndPetnames          int = 2  // ContactListAndPetnames represents NIP-02: Contact List and Petnames
	ConventionsForClients           int = 10 // 0ConventionsForClients represents NIP-10: Conventions for clients' use of `e` and `p` tags in text events
	CommandResults                  int = 20 // CommandResults represents NIP-20: Command Results
	CountingResults                 int = 45 // CountingResults represents NIP-45: Counting results
	DelegatedEventSigning           int = 26 // DelegatedEventSigning represents NIP-26: Delegated Event Signing
	EncryptedDirectMessage          int = 4  // EncryptedDirectMessage represents NIP-04: Encrypted Direct Message
	EventDeletion                   int = 9  // EventDeletion represents NIP-09: Event Deletion
	EventTreatment                  int = 16 // EventTreatment represents NIP-16: Event Treatment
	EventCreatedAtLimits            int = 22 // EventCreatedAtLimits represents NIP-22: Event `created_at` Limits
	ExternalIdentitiesInProfiles    int = 39 // ExternalIdentitiesInProfiles represents NIP-39: External Identities in Profiles
	ExpirationTimestamp             int = 40 // ExpirationTimestamp represents NIP-40: Expiration Timestamp
	FileMetadata                    int = 94 // FileMetadata represents NIP-94: File Metadata
	GenericTagQueries               int = 12 // GenericTagQueries represents NIP-12: Generic Tag Queries
	HandlingMentionsDeprecated      int = 8  // HandlingMentionsDeprecated represents NIP-08: Handling Mentions
	KeywordsFilter                  int = 50 // KeywordsFilter represents NIP-50: Keywords filter
	LightningZaps                   int = 57 // LightningZaps represents NIP-57: Lightning Zaps
	Lists                           int = 51 // Lists represents NIP-51: Lists
	LongFormContent                 int = 23 // LongFormContent represents NIP-23: Long-form Content
	MappingNostrKeysToDNS           int = 5  // MappingNostrKeysToDNS represents NIP-05: Mapping Nostr keys to DNS-based internet identifiers
	NostrMarketplace                int = 15 // NostrMarketplace represents NIP-15: Nostr Marketplace (for resilient marketplaces)
	NostrURLScheme                  int = 21 // NostrURLScheme represents NIP-21: `nostr:` URL scheme
	NostrConnect                    int = 46 // NostrConnect represents NIP-46: Nostr Connect
	OpenTimestampsAttestations      int = 3  // OpenTimestampsAttestations represents NIP-03: OpenTimestamps Attestations for Events
	ParameterizedReplaceableEvents  int = 33 // ParameterizedReplaceableEvents represents NIP-33: Parameterized Replaceable Events
	ProofOfWork                     int = 13 // ProofOfWork represents NIP-13: Proof of Work
	PublicChat                      int = 28 // PublicChat represents NIP-28: Public Chat
	Reactions                       int = 25 // Reactions represents NIP-25: Reactions
	RelayInformationDocument        int = 11 // RelayInformationDocument represents NIP-11: Relay Information Document
	RelayListMetadata               int = 65 // RelayListMetadata represents NIP-65: Relay List Metadata
	Reporting                       int = 56 // Reporting represents NIP-56: Reporting
	Reposts                         int = 18 // Reposts represents NIP-18: Reposts
	SensitiveContent                int = 36 // SensitiveContent represents NIP-36: Sensitive Content
	SubjectTagInTextEvents          int = 14 // SubjectTagInTextEvents represents NIP-14: Subject tag in text events.
	TextNoteReferences              int = 27 // TextNoteReferences represents NIP-27: Text Note References
	WindowNostrCapability           int = 7  // WindowNostrCapability represents NIP-07: `window.nostr` capability for web browsers
)
