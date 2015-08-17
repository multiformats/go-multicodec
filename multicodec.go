package multicodec

import ()

// Multicodec is the interface for a multicodec
type Multicodec interface {
	Codec

	Header() []byte
}
