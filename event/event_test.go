package event_test

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/go-nostr/nostr/event"
	"github.com/go-nostr/nostr/tag"
)

func Test_New(t *testing.T) {
	type args struct {
		Kind    int
		Content string
		Tags    []tag.Tag
	}
	tests := []struct {
		name   string
		args   args
		expect *event.Event
	}{
		{
			name: "SHOULD create new Event",
			args: args{
				Kind:    0,
				Content: "",
				Tags:    make([]tag.Tag, 0),
			},
			expect: &event.Event{
				Kind:    0,
				Content: "",
				Tags:    make([]tag.Tag, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := event.New(tt.args.Kind, tt.args.Content, tt.args.Tags...)
			if !reflect.DeepEqual(tt.expect, got) {
				t.Errorf("expected %v, got %v", tt.expect, got)
			}
			t.Logf("got %v", got)
		})
	}
}

func TestEvent_Marshal(t *testing.T) {
	type fields struct {
		event *event.Event
	}
	tests := []struct {
		name   string
		fields fields
		expect []byte
		err    error
	}{
		{
			name: "SHOULD create new Event",
			fields: fields{
				event: &event.Event{},
			},
			expect: []byte(`{}`),
			err:    nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.event.Marshal()
			if tt.err != err {
				t.Errorf("expected err %v, got err %v", tt.err, err)
			}
			if !reflect.DeepEqual(tt.expect, got) {
				t.Errorf("expected %s, got %s", tt.expect, got)
			}
			t.Logf("got %s", got)
		})
	}
}

func TestEvent_Serialize(t *testing.T) {
	type fields struct {
		event *event.Event
	}
	tests := []struct {
		name   string
		fields fields
		expect []byte
	}{
		{
			name: "SHOULD serialize Event",
			fields: fields{
				event: event.New(1, "content"),
			},
			expect: []byte(`[0,,0,1,[],content]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.event.Serialize()
			if !reflect.DeepEqual(tt.expect, got) {
				t.Errorf("expected %s, got %s", tt.expect, got)
			}
			t.Logf("got %s", got)
		})
	}
}

func TestEvent_Sign(t *testing.T) {
	prvKey, err := btcec.NewPrivateKey()
	if err != nil {
		t.Error(err)
	}
	type args struct {
		prvKeyHex string
	}
	type fields struct {
		event *event.Event
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		err    error
	}{
		{
			name: "SHOULD sign Event",
			args: args{
				prvKeyHex: hex.EncodeToString(prvKey.Serialize()),
			},
			fields: fields{
				event: event.New(1, "content"),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fields.event.Sign(tt.args.prvKeyHex)
			if err != nil {
				t.Errorf("expected %s, got %s", tt.err, err)
			}
			t.Logf("got %s", err)
		})
	}
}

func TestEvent_Verify(t *testing.T) {
	prvKey, err := btcec.NewPrivateKey()
	if err != nil {
		t.Error(err)
	}
	type fields struct {
		event *event.Event
	}
	tests := []struct {
		name    string
		fields  fields
		err     error
		wantErr bool
	}{
		{
			name: "SHOULD serialize Event",
			fields: fields{
				event: (func() *event.Event {
					evnt := event.New(1, "content")
					err := evnt.Sign(hex.EncodeToString(prvKey.Serialize()))
					if err != nil {
						t.Error(err)
					}
					return evnt
				})(),
			},
			err:     nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fields.event.Verify()
			if err != nil {
				t.Error(err)
			}
			t.Logf("got %s", err)
		})
	}
}
