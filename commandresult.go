package nostr

// CommandResult TODO
type CommandResult struct {
	Status  []byte `json:"status,omitempty"`
	EventID []byte `json:"event_id,omitempty"`
	IsSaved bool   `json:"is_saved,omitempty"`
	Message []byte `json:"message,omitempty"`
}
