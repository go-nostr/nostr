package subscriptionid_test

import (
	"fmt"
	"testing"

	"github.com/go-nostr/nostr/subscriptionid"
)

func Test_New(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "SHOULD construct non-zero string value of length 64",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subscriptionid.New()
			if len(got) != 64 {
				t.Fatalf("expected non-zero string value of length 64")
			}
			t.Logf("got %v", got)
		})
	}
}

func Test_Validate(t *testing.T) {
	type args struct {
		subscriptionID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "SHOULD NOT validate non-zero string value greater than 64 characters in length",
			args: args{
				subscriptionID: fmt.Sprintf("%s%s", subscriptionid.New(), "abcd-1234"),
			},
			wantErr: true,
		},
		{
			name: "SHOULD validate non-zero string value less than or equal to 64 characters in length",
			args: args{
				subscriptionID: subscriptionid.New(),
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := subscriptionid.Validate(tt.args.subscriptionID)
			if got != nil && !tt.wantErr {
				t.Fatalf("expected nil, got %v", got)
			}
			t.Logf("got %v", got)
		})
	}
}
