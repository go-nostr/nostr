package okmessage_test

import (
	"reflect"
	"testing"

	"github.com/go-nostr/nostr/message"
	"github.com/go-nostr/nostr/message/okmessage"
)

func Test_New(t *testing.T) {
	type args struct {
		id     string
		ok     bool
		status string
	}
	tests := []struct {
		name   string
		args   args
		expect message.Message
	}{
		{
			name: "SHOULD create new message.Message",
			args: args{
				id:     "asdf",
				ok:     true,
				status: "asdfasdf",
			},
			expect: message.Message{"OK", "asdf", true, "asdfasdf"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := okmessage.New(tt.args.id, tt.args.ok, tt.args.status)
			if !reflect.DeepEqual(tt.expect, got) {
				t.Errorf("expected %v, got %v", tt.expect, got)
				return
			}
			t.Logf("got %v", got)
		})
	}
}
