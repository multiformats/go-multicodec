package multicodec

import (
	"bytes"
	"errors"
	"io"
)

var (
	ErrHeaderInvalid = errors.New("multicodec header invalid")
	ErrMismatch      = errors.New("multicodec did not match")
	ErrVarints       = errors.New("multicodec varints not yet implemented")
)

// Header returns a multicodec header with the given path.
func Header(path []byte) []byte {
	l := len(path)
	if l >= 127 {
		panic(ErrVarints.Error())
	}

	buf := make([]byte, l+2)
	buf[0] = byte(l)
	copy(buf[1:], path)
	buf[l+1] = '\n'
	return buf
}

// WriteHeader writes a multicodec header to a writer.
// It uses the given path.
func WriteHeader(w io.Writer, path []byte) error {
	hdr := Header(path)
	_, err := w.Write(hdr)
	return err
}

// ReadHeader reads a multicodec header from a reader.
// Returns the path found, or an error if the header
// mismatched.
func ReadHeader(r io.Reader) (path []byte, err error) {

	buf := make([]byte, 1)
	if _, err := r.Read(buf); err != nil {
		return nil, err
	}

	l := int(buf[0])
	if l > 127 {
		return nil, ErrVarints
	}

	buf = make([]byte, l+1)
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}
	if buf[l] != '\n' {
		return nil, ErrHeaderInvalid
	}
	return buf[:l], nil
}

// ConsumePath reads a multicodec header from a reader,
// verifying it matches given path. If it does not, it returns
// ErrProtocolMismatch
func ConsumePath(r io.Reader, path []byte) (err error) {
	actual, err := ReadHeader(r)
	if !bytes.Equal(path, actual) {
		return ErrMismatch
	}
	return nil
}

// ConsumeHeader reads a multicodec header from a reader,
// verifying it matches given header. If it does not, it returns
// ErrProtocolMismatch
func ConsumeHeader(r io.Reader, header []byte) (err error) {
	actual := make([]byte, len(header))
	if _, err := io.ReadFull(r, actual); err != nil {
		return err
	}

	if !bytes.Equal(header, actual) {
		return ErrMismatch
	}
	return nil
}
