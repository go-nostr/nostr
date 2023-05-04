package authmessage_test

import (
	"reflect"
	"testing"

	"github.com/go-nostr/nostr/event"
	"github.com/go-nostr/nostr/message"
	"github.com/go-nostr/nostr/message/authmessage"
)

func Test_New(t *testing.T) {
	type args struct {
		opt *authmessage.Options
	}
	tests := []struct {
		name   string
		args   args
		expect *message.Message
	}{
		{
			name: "MUST create new auth message with challenge",
			args: args{
				opt: &authmessage.Options{
					Challenge: "challenge",
				},
			},
			expect: &message.Message{authmessage.Type, "challenge"},
		},
		{
			name: "MUST create new auth message with event",
			args: args{
				opt: &authmessage.Options{
					Event: event.New(0, ""),
				},
			},
			expect: &message.Message{authmessage.Type, event.New(0, "")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := authmessage.New(tt.args.opt)
			if !reflect.DeepEqual(tt.expect, got) {
				t.Errorf("expected %v, got %v", tt.expect, got)
			}
			t.Logf("got %v", got)
		})
	}
}
