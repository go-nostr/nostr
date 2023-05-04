package npub_test

import (
	"testing"

	"github.com/go-nostr/nostr/npub"
)

func Test_Decode(t *testing.T) {
	type args struct {
		npub string
	}
	tests := []struct {
		name   string
		args   args
		expect string
	}{
		{
			name: "SHOULD decode npub",
			args: args{
				npub: "npub10elfcs4fr0l0r8af98jlmgdh9c8tcxjvz9qkw038js35mp4dma8qzvjptg",
			},
			expect: "7e7e9c42a91bfef19fa929e5fda1b72e0ebc1a4c1141673e2794234d86addf4e",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := npub.Decode(tt.args.npub)
			if err != nil {
				t.Errorf("%v", err.Error())
			}
			if tt.expect != got {
				t.Errorf("expected %v, got %v", tt.expect, got)
			}
			t.Logf("got %v", got)
		})
	}
}

func Test_Encode(t *testing.T) {
	type args struct {
		pubKeyHex string
	}
	tests := []struct {
		name   string
		args   args
		expect string
	}{
		{
			name: "SHOULD decode npub",
			args: args{
				pubKeyHex: "7e7e9c42a91bfef19fa929e5fda1b72e0ebc1a4c1141673e2794234d86addf4e",
			},
			expect: "npub10elfcs4fr0l0r8af98jlmgdh9c8tcxjvz9qkw038js35mp4dma8qzvjptg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := npub.Encode(tt.args.pubKeyHex)
			if err != nil {
				t.Errorf("%v", err.Error())
			}
			if tt.expect != got {
				t.Errorf("expected %v, got %v", tt.expect, got)
			}
			t.Logf("got %v", got)
		})
	}
}
