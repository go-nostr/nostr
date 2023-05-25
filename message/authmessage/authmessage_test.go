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
		challenge string
		evt       *event.Event
	}
	tests := []struct {
		name string
		args args
		want message.Message
	}{
		{
			name: "MUST create new auth message with challenge",
			args: args{
				challenge: "challenge",
			},
			want: message.Message{authmessage.Type, "challenge"},
		},
		{
			name: "MUST create new auth message with event",
			args: args{
				evt: event.New(0, ""),
			},
			want: message.Message{authmessage.Type, event.New(0, "")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := authmessage.New(tt.args.challenge, tt.args.evt)
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("wanted %v, got %v", tt.want, got)
			}
			t.Logf("got %v", got)
		})
	}
}
