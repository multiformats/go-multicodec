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
	l := len(path) + 1 // + \n
	if l >= 127 {
		panic(ErrVarints.Error())
	}

	buf := make([]byte, l+1)
	buf[0] = byte(l)
	copy(buf[1:], path)
	buf[l] = '\n'
	return buf
}

// HeaderPath returns the multicodec path from header
func HeaderPath(hdr []byte) []byte {
	hdr = hdr[1:]
	if hdr[len(hdr)-1] == '\n' {
		hdr = hdr[:len(hdr)-1]
	}
	return hdr
}

// WriteHeader writes a multicodec header to a writer.
// It uses the given path.
func WriteHeader(w io.Writer, path []byte) error {
	hdr := Header(path)
	_, err := w.Write(hdr)
	return err
}

// ReadHeader reads a multicodec header from a reader.
// Returns the header found, or an error if the header
// mismatched.
func ReadHeader(r io.Reader) (path []byte, err error) {
	lbuf := make([]byte, 1)
	if _, err := r.Read(lbuf); err != nil {
		return nil, err
	}

	l := int(lbuf[0])
	if l > 127 {
		return nil, ErrVarints
	}

	buf := make([]byte, l+1)
	buf[0] = lbuf[0]
	if _, err := io.ReadFull(r, buf[1:]); err != nil {
		return nil, err
	}
	if buf[l] != '\n' {
		return nil, ErrHeaderInvalid
	}
	return buf, nil
}

// ReadPath reads a multicodec header from a reader.
// Returns the path found, or an error if the header
// mismatched.
func ReadPath(r io.Reader) (path []byte, err error) {
	hdr, err := ReadHeader(r)
	if err != nil {
		return nil, err
	}
	return HeaderPath(hdr), nil
}

// ConsumePath reads a multicodec header from a reader,
// verifying it matches given path. If it does not, it returns
// ErrProtocolMismatch
func ConsumePath(r io.Reader, path []byte) (err error) {
	actual, err := ReadPath(r)
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

// WrapHeaderReader returns a reader that first reads the
// given header, and then the given reader, using io.MultiReader.
// It is useful if the header has been read through, but still
// needed to pass to a decoder.
func WrapHeaderReader(hdr []byte, r io.Reader) io.Reader {
	return io.MultiReader(bytes.NewReader(hdr), r)
}
