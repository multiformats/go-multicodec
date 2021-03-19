// Package multicodec exposes the multicodec table as Go constants.
package multicodec

//go:generate go run gen.go
//go:generate gofmt -w code.go
//go:generate go get golang.org/x/tools/cmd/stringer
//go:generate go run golang.org/x/tools/cmd/stringer -type=Code -linecomment
//go:generate go mod tidy
