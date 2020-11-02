// Package multicodec exposes the table of multicodecs as Go constants.
package multicodec

//go:generate go run gen.go
//go:generate gofmt -w codec.go
//go:generate stringer -type=Codec -linecomment
