package message_test

import (
	"reflect"
	"testing"

	"github.com/go-nostr/nostr/message"
)

func Test_New(t *testing.T) {
	tests := []struct {
		name   string
		expect message.Message
	}{
		{
			name:   "MUST create new message",
			expect: message.Message{"type"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := message.New("type")
			if !reflect.DeepEqual(tt.expect, got) {
				t.Fatalf("expected %v, got %v", tt.expect, got)
			}
			t.Logf("got %v", got)
		})
	}
}

func TestMessage_Marshal(t *testing.T) {
	type args struct{}
	type fields struct {
		mess message.Message
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		expect []byte
		err    error
	}{
		{
			name: "MUST marshal message",
			args: args{},
			fields: fields{
				mess: message.New("type"),
			},
			expect: []byte("[\"type\"]"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mess, err := tt.fields.mess.Marshal()
			if err != nil && tt.err == nil {
				t.Fatalf("expected error %v, got error %v", tt.err, err)
			}
			if !reflect.DeepEqual(mess, tt.expect) {
				t.Fatalf("expected %v, got %v", tt.expect, mess)
			}
			t.Logf("got %+v", mess)
		})
	}
}

func TestMessage_Push(t *testing.T) {
	type args struct {
		elem any
	}
	type fields struct {
		mess message.Message
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		expect message.Message
	}{
		{
			name: "MUST push value into message",
			args: args{
				elem: "type",
			},
			fields: fields{
				mess: message.New(),
			},
			expect: message.Message{"type"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.mess.Push(tt.args.elem)
			if !reflect.DeepEqual(tt.expect, tt.fields.mess) {
				t.Fatalf("expected %v, got %v", tt.expect, tt.fields.mess)
			}
			t.Logf("got %v", tt.fields.mess)
		})
	}
}

func TestMessage_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
	}
	type fields struct {
		mess message.Message
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		expect message.Message
		err    error
	}{
		{
			name: "MUST unmarshal message",
			args: args{
				data: []byte("[\"type\"]"),
			},
			fields: fields{
				mess: message.Message{},
			},
			expect: message.New("type"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fields.mess.Unmarshal(tt.args.data)
			if tt.err != nil && err == nil {
				t.Fatalf("expected error %v, got error %v", tt.err, err)
			}
			if !reflect.DeepEqual(tt.fields.mess, tt.expect) {
				t.Fatalf("expected %v, got %v", tt.expect, tt.fields.mess)
			}
			t.Logf("got %+v", tt.fields.mess)
		})
	}
}

func TestMessage_Values(t *testing.T) {
	type fields struct {
		mess message.Message
	}
	tests := []struct {
		name   string
		fields fields
		expect []any
		err    error
	}{
		{
			name: "MUST get values for raw message",
			fields: fields{
				mess: message.New("type"),
			},
			expect: []any{"type"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fields.mess.Values()
			if !reflect.DeepEqual(tt.expect, got) {
				t.Fatalf("expected %v, got %v", tt.expect, got)
			}
			t.Logf("got %v", got)
		})
	}
}
