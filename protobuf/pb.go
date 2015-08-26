package mc_cbor

import (
	"errors"
	"io"

	proto "github.com/golang/protobuf/proto"
	msgio "github.com/jbenet/go-msgio"

	mc "github.com/jbenet/go-multicodec"
)

var Header []byte

var ErrNotProtobuf = errors.New("not a protobuf")

func init() {
	Header = mc.Header([]byte("/protobuf/msgio"))
}

type codec struct {
	mc bool
}

func Codec(m proto.Message) mc.Codec {
	return &codec{mc: false}
}

func Multicodec(m proto.Message) mc.Multicodec {
	return &codec{mc: true}
}

func (c *codec) Encoder(w io.Writer) mc.Encoder {
	return &encoder{
		rawW: w,
		w:    msgio.NewWriter(w),
		buf:  proto.NewBuffer(nil),
		mc:   c.mc,
	}
}

func (c *codec) Decoder(r io.Reader) mc.Decoder {
	return &decoder{
		rawR: r,
		r:    msgio.NewReader(r),
		mc:   c.mc,
	}
}

func (c *codec) Header() []byte {
	return Header
}

type encoder struct {
	rawW io.Writer
	w    msgio.Writer
	buf  *proto.Buffer
	mc   bool
}

type decoder struct {
	rawR io.Reader
	r    msgio.Reader
	mc   bool
}

func (c *encoder) Encode(v interface{}) error {
	// if multicodec, write the header first
	if c.mc {
		if _, err := c.rawW.Write(Header); err != nil {
			return err
		}
	}

	pbv, ok := v.(proto.Message)
	if !ok {
		return ErrNotProtobuf
	}

	defer c.buf.Reset()
	if err := c.buf.Marshal(pbv); err != nil {
		return err
	}
	return c.w.WriteMsg(c.buf.Bytes())
}

func (c *decoder) Decode(v interface{}) error {
	// if multicodec, consume the header first
	if c.mc {
		if err := mc.ConsumeHeader(c.rawR, Header); err != nil {
			return err
		}
	}

	pbv, ok := v.(proto.Message)
	if !ok {
		return ErrNotProtobuf
	}

	buf, err := c.r.ReadMsg()
	if err != nil {
		return err
	}

	return proto.Unmarshal(buf, pbv)
}
