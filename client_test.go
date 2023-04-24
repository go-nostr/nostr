package nostr_test

import (
	"testing"

	"github.com/go-nostr/nostr"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "SHOULD create client",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cl := nostr.NewClient()

			if cl == nil {
				t.Fatalf("expected %+v, to be not nil", cl)
			}

			t.Log(cl)
		})
	}
}
