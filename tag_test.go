package nostr_test

import (
	"reflect"
	"testing"

	"github.com/go-nostr/nostr"
)

func Test_NewAmountTag(t *testing.T) {
	type args struct {
		amount int
	}
	tests := []struct {
		name   string
		args   args
		expect *nostr.AmountTag
	}{
		{
			name: "SHOULD create new AmountTag",
			args: args{
				amount: 42,
			},
			expect: &nostr.AmountTag{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag := nostr.NewAmountTag(tt.args.amount)
			if reflect.DeepEqual(tt.expect, tag) {
				t.Fatalf("expected %v, got %v", tt.expect, tag)
			}
			t.Logf("%+v", tag)
		})
	}
}

func TestAmountTag_Marshal(t *testing.T) {
	type fields struct {
		amount int
	}
	tests := []struct {
		name   string
		fields fields
		expect []byte
	}{
		{
			name: "SHOULD marshal AmountTag",
			fields: fields{
				amount: 42,
			},
			expect: []byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amountTag := nostr.NewAmountTag(tt.fields.amount)
			data, err := amountTag.Marshal()
			if err != nil {
				t.Errorf("%+v", err)
			}
			t.Logf("%+v", data)
		})
	}
}

func TestAmountTag_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
	}
	type fields struct {
		amount int
	}
	tests := []struct {
		name   string
		args   args
		fields fields
		expect *nostr.AmountTag
	}{
		{
			name: "SHOULD unmarshal AmountTag",
			args: args{
				data: []byte("[\"amount\",42]"),
			},
			fields: fields{
				amount: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			amountTag := &nostr.AmountTag{}
			err := amountTag.Unmarshal(tt.args.data)
			if err != nil {
				t.Errorf("%+v", err)
			}
			if reflect.DeepEqual(amountTag, tt.expect) {
				t.Errorf("expected %+v, got %+v", amountTag, tt.expect)
			}
		})
	}
}
