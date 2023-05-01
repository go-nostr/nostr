package nostr

// Limitations represents various limits and constraints imposed by a Nostr server.
type Limitations struct {
	MaxMessageLength int64 `json:"max_message_length,omitempty"` // Maximum length of a message.
	MaxSubscriptions int64 `json:"max_subscriptions,omitempty"`  // Maximum number of subscriptions.
	MaxFilters       int64 `json:"max_filters,omitempty"`        // Maximum number of filters.
	MaxLimit         int64 `json:"max_limit,omitempty"`          // Maximum limit for query results.
	MaxSubIDLength   int64 `json:"max_subid_length,omitempty"`   // Maximum length of a subscription ID.
	MinPrefix        int64 `json:"min_prefix,omitempty"`         // Minimum prefix length.
	MaxEventTags     int64 `json:"max_event_tags,omitempty"`     // Maximum number of event tags.
	MaxContentLength int64 `json:"max_content_length,omitempty"` // Maximum content length.
	MinPowDifficulty int64 `json:"min_pow_difficulty,omitempty"` // Minimum proof-of-work difficulty.
	AuthRequired     bool  `json:"auth_required,omitempty"`      // Indicates if authentication is required.
	PaymentRequired  bool  `json:"payment_required,omitempty"`   // Indicates if payment is required.
}
