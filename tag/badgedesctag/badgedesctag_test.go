package badgedesctag_test

import (
	"reflect"
	"testing"

	"github.com/go-nostr/nostr/tag/badgedesctag"
)

func Test_New(t *testing.T) {
	type args struct {
		desc string
	}
	tests := []struct {
		name   string
		args   args
		expect []any
	}{
		{
			name: "SHOULD create new authmessage.Message",
			args: args{
				desc: "This is only a test.",
			},
			expect: []any{badgedesctag.Type, "This is only a test."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := badgedesctag.New(tt.args.desc)
			if !reflect.DeepEqual(tt.expect, got.Values()) {
				t.Errorf("expected %v, got %v", tt.expect, got.Values())
				return
			}
			t.Logf("got %v", got.Values()...)
		})
	}
}
