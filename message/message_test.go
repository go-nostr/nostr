package message_test

import (
	"reflect"
	"testing"

	"github.com/go-nostr/nostr/message"
)

func Test_NewMessage(t *testing.T) {
	tests := []struct {
		name   string
		expect *message.Message
	}{
		{
			name:   "MUST create a new Message with given type",
			expect: &message.Message{[]byte("\"type\"")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mess := message.New("type")
			if !reflect.DeepEqual(mess, tt.expect) {
				t.Fatalf("expected %v, got %v", tt.expect, mess)
			}
			t.Logf("got %+v", mess)
		})
	}
}

func TestMessage_Marshal(t *testing.T) {
	type args struct{}
	type fields struct {
		mess *message.Message
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		expect []byte
		err    error
	}{
		{
			name: "MUST create a new Message with given type",
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

func TestMessage_Type(t *testing.T) {
	type fields struct {
		mess *message.Message
	}
	tests := []struct {
		name   string
		fields fields
		expect []byte
		err    error
	}{
		{
			name: "MUST get type for raw message",
			fields: fields{
				mess: message.New("type"),
			},
			expect: []byte("type"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			typ := tt.fields.mess.Type()
			if !reflect.DeepEqual(typ, tt.expect) {
				t.Fatalf("expected %v, got %v", tt.expect, tt.fields.mess)
			}
			t.Logf("got %+v", tt.fields.mess)
		})
	}
}

func TestMessage_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
	}
	type fields struct {
		mess *message.Message
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		expect *message.Message
		err    error
	}{
		{
			name: "MUST unmarshal Message",
			args: args{
				data: []byte("[\"type\"]"),
			},
			fields: fields{
				mess: &message.Message{},
			},
			expect: message.New("type"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.fields.mess.Unmarshal(tt.args.data)
			if err != nil && tt.err == nil {
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
		mess *message.Message
	}
	tests := []struct {
		name   string
		fields fields
		expect []any
		err    error
	}{
		{
			name: "MUST get type for raw message",
			fields: fields{
				mess: message.New("type"),
			},
			expect: []any{"type"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := tt.fields.mess.Values()
			if !reflect.DeepEqual(v, tt.expect) {
				t.Fatalf("expected %v, got %v", tt.expect, tt.fields.mess)
			}
			t.Logf("got %+v", tt.fields.mess)
		})
	}
}
