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

func Test_New(t *testing.T) {
	tests := []struct {
		name   string
		expect string
	}{
		{
			name: "SHOULD decode npub",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prvKeyHex, pubKeyHex, npubStr, err := npub.New()
			if err != nil {
				t.Errorf("%v", err.Error())
			}
			if npubStr == "" {
				t.Errorf("expected npub != \"\"")
			}
			t.Logf("got npub: %v", npubStr)
			if pubKeyHex == "" {
				t.Errorf("expected pubKeyHex != \"\"")
			}
			t.Logf("got pubKeyHex: %v", pubKeyHex)
			if prvKeyHex == "" {
				t.Errorf("expected prvKeyHex != \"\"")
			}
			t.Logf("got prvKeyHex: %v", prvKeyHex)
			npubStr2, err := npub.Encode(pubKeyHex)
			if err != nil {
				t.Error(err)
			}
			if npubStr != npubStr2 {
				t.Errorf("expected npubStr %v, got npubStr %v", npubStr, npubStr2)
			}
			t.Logf("got npubstr2: %v", npubStr2)
			pubKeyHex2, err := npub.Decode(npubStr)
			if err != nil {
				t.Error(err)
			}
			if pubKeyHex != pubKeyHex2 {
				t.Errorf("expected pubKeyHex %v, got pubKeyHex %v", pubKeyHex, pubKeyHex2)
			}
			t.Logf("got pubKeyHex2: %v", pubKeyHex2)
		})
	}
}
