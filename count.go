package nostr

// Count represents the count of a specific item or event.
type Count struct {
	Count int `json:"count,omitempty"` // The number of occurrences of the item or event.
}
