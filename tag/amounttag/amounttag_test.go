package amounttag_test

import (
	"reflect"
	"testing"

	"github.com/go-nostr/nostr/tag"
	"github.com/go-nostr/nostr/tag/amounttag"
)

func Test_New(t *testing.T) {
	type args struct {
		amount int
	}
	tests := []struct {
		name   string
		args   args
		expect *tag.Tag
	}{
		{
			name: "SHOULD create new tag.Tag",
			args: args{
				amount: 1984,
			},
			expect: &tag.Tag{"amount", 1984},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := amounttag.New(tt.args.amount)
			if !reflect.DeepEqual(tt.expect, got) {
				t.Errorf("expected %v, got %v", tt.expect, got)
				return
			}
			t.Logf("got %v", got)
		})
	}
}

func Test_Type(t *testing.T) {
	tests := []struct {
		name   string
		expect string
	}{
		{
			name:   "SHOULD expect type to equal 'amount'",
			expect: "amount",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := amounttag.Type
			if !reflect.DeepEqual(tt.expect, got) {
				t.Errorf("expected %v, got %v", tt.expect, got)
				return
			}
			t.Logf("got %v", got)
		})
	}
}
