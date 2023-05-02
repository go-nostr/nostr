package nostr

import "github.com/go-nostr/nostr/tag"

// InformationDocument contains information about a Nostr relay.
type InformationDocument struct {
	Name           string       `json:"name,omitempty"`
	Description    string       `json:"description,omitempty"`
	PubKey         string       `json:"pub_key,omitempty"`
	Contact        string       `json:"contact,omitempty"`
	SupportedNIPs  string       `json:"supported_nips,omitempty"`
	Software       string       `json:"software,omitempty"`
	Version        string       `json:"version,omitempty"`
	Limitations    *Limitations `json:"limitations,omitempty"`
	LanguageTags   []string     `json:"language_tags,omitempty"`
	RelayCountries []string     `json:"relay_countries,omitempty"`
	Tags           []tag.Tag    `json:"tags,omitempty"`
	PostingPolicy  string       `json:"posting_policy,omitempty"`
	PaymentsURL    string       `json:"payments_url,omitempty"`
	Fees           *Fees        `json:"fees,omitempty"`
}
