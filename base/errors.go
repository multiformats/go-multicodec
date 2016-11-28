package base

import "errors"

var (
	ErrExpectedReader = errors.New("expected io.Reader as input")
	ErrExpectedWriter = errors.New("expected io.Writer as input")
)
