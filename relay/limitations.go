package relay

// Limitations represents various limits and constraints imposed by a Nostr server.
type Limitations struct {
	AuthRequired     bool `json:"auth_required,omitempty"`      // Indicates if authentication is required.
	MaxContentLength int  `json:"max_content_length,omitempty"` // Maximum content length.
	MaxEventTags     int  `json:"max_event_tags,omitempty"`     // Maximum number of event tags.
	MaxFilters       int  `json:"max_filters,omitempty"`        // Maximum number of filters.
	MaxLimit         int  `json:"max_limit,omitempty"`          // Maximum limit for query results.
	MaxMessageLength int  `json:"max_message_length,omitempty"` // Maximum length of a message.
	MaxSubIDLength   int  `json:"max_subid_length,omitempty"`   // Maximum length of a subscription ID.
	MaxSubscriptions int  `json:"max_subscriptions,omitempty"`  // Maximum number of subscriptions.
	MinPowDifficulty int  `json:"min_pow_difficulty,omitempty"` // Minimum proof-of-work difficulty.
	MinPrefix        int  `json:"min_prefix,omitempty"`         // Minimum prefix length.
	PaymentRequired  bool `json:"payment_required,omitempty"`   // Indicates if payment is required.
}
