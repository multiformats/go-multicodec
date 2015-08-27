package mc_json

import (
	"encoding/json"
	"io"

	mc "github.com/jbenet/go-multicodec"
)

var Header []byte

func init() {
	Header = mc.Header([]byte("/json"))
}

type codec struct {
	mc bool
}

func Codec() mc.Codec {
	return &codec{mc: false}
}

func Multicodec() mc.Multicodec {
	return &codec{mc: true}
}

func (c *codec) Encoder(w io.Writer) mc.Encoder {
	return &encoder{
		w:   w,
		mc:  c.mc,
		enc: json.NewEncoder(w),
	}
}

func (c *codec) Decoder(r io.Reader) mc.Decoder {
	return &decoder{
		r:   r,
		mc:  c.mc,
		dec: json.NewDecoder(r),
	}
}

func (c *codec) Header() []byte {
	return Header
}

type encoder struct {
	w   io.Writer
	mc  bool
	enc *json.Encoder
}

type decoder struct {
	r   io.Reader
	mc  bool
	dec *json.Decoder
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
