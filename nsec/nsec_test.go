package nsec_test

import (
	"testing"

	"github.com/go-nostr/nostr/nsec"
)

func Test_Decode(t *testing.T) {
	type args struct {
		nsec string
	}
	tests := []struct {
		name   string
		args   args
		expect string
	}{
		{
			name: "SHOULD decode nsec",
			args: args{
				nsec: "nsec1vl029mgpspedva04g90vltkh6fvh240zqtv9k0t9af8935ke9laqsnlfe5",
			},
			expect: "67dea2ed018072d675f5415ecfaed7d2597555e202d85b3d65ea4e58d2d92ffa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nsec.Decode(tt.args.nsec)
			if err != nil {
				t.Errorf("%v", err.Error())
			}
			if tt.expect != got {
				t.Errorf("expected %v, got %v", tt.expect, got)
			}
			t.Logf("got %v", got)
		})
	}
}

func Test_Encode(t *testing.T) {
	type args struct {
		pubKeyHex string
	}
	tests := []struct {
		name   string
		args   args
		expect string
	}{
		{
			name: "SHOULD decode nsec",
			args: args{
				pubKeyHex: "67dea2ed018072d675f5415ecfaed7d2597555e202d85b3d65ea4e58d2d92ffa",
			},
			expect: "nsec1vl029mgpspedva04g90vltkh6fvh240zqtv9k0t9af8935ke9laqsnlfe5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := nsec.Encode(tt.args.pubKeyHex)
			if err != nil {
				t.Errorf("%v", err.Error())
			}
			if tt.expect != got {
				t.Errorf("expected %v, got %v", tt.expect, got)
			}
			t.Logf("got %v", got)
		})
	}
}
