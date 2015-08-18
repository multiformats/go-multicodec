package mc_json

import (
	"io"

	ugcodec "github.com/ugorji/go/codec"

	mc "github.com/jbenet/go-multicodec"
)

var Header []byte
var DefaultHandle *ugcodec.JsonHandle

func init() {
	Header = mc.Header([]byte("/json"))
	DefaultHandle = &ugcodec.JsonHandle{}
	DefaultHandle.Canonical = true
}

type codec struct {
	Handle *ugcodec.JsonHandle
	mc     bool
}

func Codec(h *ugcodec.JsonHandle) mc.Codec {
	if h == nil {
		h = DefaultHandle
	}
	return &codec{
		Handle: h,
		mc:     false,
	}
}

func Multicodec(h *ugcodec.JsonHandle) mc.Multicodec {
	if h == nil {
		h = DefaultHandle
	}
	return &codec{
		Handle: h,
		mc:     true,
	}
}

func (c *codec) Encoder(w io.Writer) mc.Encoder {
	return &encoder{
		w:   w,
		mc:  c.mc,
		enc: ugcodec.NewEncoder(w, c.Handle),
	}
}

func (c *codec) Decoder(r io.Reader) mc.Decoder {
	return &decoder{
		r:   r,
		mc:  c.mc,
		dec: ugcodec.NewDecoder(r, c.Handle),
	}
}

func (c *codec) Header() []byte {
	return Header
}

type encoder struct {
	w   io.Writer
	mc  bool
	enc *ugcodec.Encoder
}

type decoder struct {
	r   io.Reader
	mc  bool
	dec *ugcodec.Decoder
}

func (c *encoder) Encode(v interface{}) error {
	// if multicodec, write the header first
	if c.mc {
		if _, err := c.w.Write(Header); err != nil {
			return err
		}
	}
	return c.enc.Encode(v)
}

func (c *decoder) Decode(v interface{}) error {
	// if multicodec, consume the header first
	if c.mc {
		if err := mc.ConsumeHeader(c.r, Header); err != nil {
			return err
		}
	}
	return c.dec.Decode(v)
}
