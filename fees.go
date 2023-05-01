package nostr

// Fees is a struct that contains information about the fees associated with an operation.
type Fees struct {
	Admission *Admission `json:"admission,omitempty"` // Admission represents the admission fee for an operation, if applicable.
}
