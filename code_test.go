package multicodec_test

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/multiformats/go-multicodec"
)

func TestFlagValue(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		in   string

		want    multicodec.Code
		wantErr string
	}{
		{
			name: "EmptyString",
			in:   "",

			wantErr: "unknown multicodec",
		},
		{
			name: "Name",
			in:   "sha1",

			want: multicodec.Sha1,
		},
		{
			name: "NumberHex",
			in:   "0x11",

			want: multicodec.Sha1,
		},
		{
			name: "NumberDecimal",
			in:   "17",

			want: multicodec.Sha1,
		},
		{
			name: "ReservedNumber",
			in:   "0x300012",

			want: multicodec.Code(0x300012),
		},

		{
			name: "UnknownName",
			in:   "some-name",

			wantErr: "unknown multicodec",
		},
		{
			name: "UnknownNumber",
			in:   "0x1000",

			wantErr: "unknown multicodec",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			fs := flag.NewFlagSet("", flag.ContinueOnError)
			fs.SetOutput(ioutil.Discard)

			var code multicodec.Code
			fs.Var(&code, "multicodec", "")

			err := fs.Parse([]string{"--multicodec=" + test.in})
			if test.wantErr == "" {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			} else {
				gotErr := fmt.Sprint(err)
				if !strings.Contains(gotErr, test.wantErr) {
					t.Fatalf("want error %q, got: %v", test.wantErr, err)
				}
			}
		})
	}

}
