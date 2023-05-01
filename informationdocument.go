package nostr

// InformationDocument contains information about a Nostr relay.
type InformationDocument struct {
	Name           []byte       `json:"name,omitempty"`
	Description    []byte       `json:"description,omitempty"`
	PubKey         []byte       `json:"pub_key,omitempty"`
	Contact        []byte       `json:"contact,omitempty"`
	SupportedNIPs  []byte       `json:"supported_nips,omitempty"`
	Software       []byte       `json:"software,omitempty"`
	Version        []byte       `json:"version,omitempty"`
	Limitations    *Limitations `json:"limitations,omitempty"`
	LanguageTags   [][]byte     `json:"language_tags,omitempty"`
	RelayCountries [][]byte     `json:"relay_countries,omitempty"`
	Tags           []Tag        `json:"tags,omitempty"`
	PostingPolicy  []byte       `json:"posting_policy,omitempty"`
	PaymentsURL    []byte       `json:"payments_url,omitempty"`
	Fees           *Fees        `json:"fees,omitempty"`
}
