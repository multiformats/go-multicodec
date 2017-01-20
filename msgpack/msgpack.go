package mc_msgpack

import (
	"io"

	mc "github.com/multiformats/go-multicodec"
	gocodec "github.com/ugorji/go/codec"
)

var HeaderPath string
var Header []byte

func init() {
	HeaderPath = "/msgpack"
	Header = mc.Header([]byte(HeaderPath))
}

type codec struct {
	mc bool
}

func Codec() mc.Codec {
	return &codec{
		mc: false,
	}
}

func Multicodec() mc.Multicodec {
	return &codec{
		mc: true,
	}
}

func (c *codec) Encoder(w io.Writer) mc.Encoder {
	h := &gocodec.MsgpackHandle{}
	h.Canonical = true
	return &encoder{
		w:   w,
		mc:  c.mc,
		enc: gocodec.NewEncoder(w, h),
	}
}

func (c *codec) Decoder(r io.Reader) mc.Decoder {
	h := &gocodec.MsgpackHandle{}
	h.Canonical = true
	return &decoder{
		r:   r,
		mc:  c.mc,
		dec: gocodec.NewDecoder(r, h),
	}
}

func (c *codec) Header() []byte {
	return Header
}

type encoder struct {
	w   io.Writer
	mc  bool
	enc *gocodec.Encoder
}

type decoder struct {
	r   io.Reader
	mc  bool
	dec *gocodec.Decoder
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
