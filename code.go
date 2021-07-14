// Package multicodec exposes the multicodec table as Go constants.
package multicodec

import (
	"flag"
	"fmt"
)

//go:generate go run gen.go
//go:generate gofmt -w code_table.go
//go:generate go get golang.org/x/tools/cmd/stringer
//go:generate go run golang.org/x/tools/cmd/stringer -type=Code -linecomment
//go:generate go mod tidy

// Code describes an integer reserved in the multicodec table, defined at
// github.com/multiformats/multicodec.
type Code uint64

// Assert that Code implements flag.Value.
// Requires a pointer, since Set modifies the receiver.
var _ flag.Value = (*Code)(nil)

// Assert that Code implements fmt.Stringer without a pointer.
var _ fmt.Stringer = Code(0)

func (c *Code) Set(text string) error {
	for code, name := range _Code_map {
		if name == text {
			*c = code
			return nil
		}
	}
	return fmt.Errorf("unknown multicodec: %q", text)
}
