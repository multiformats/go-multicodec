package bin

import (
	"encoding/hex"
	"io"

	mc "github.com/multiformats/go-multicodec"
	base "github.com/multiformats/go-multicodec/base"
)

var (
	HeaderPath = "/base16/"
	Header     = mc.Header([]byte(HeaderPath))
	multic     = mc.NewMulticodecFromCodec(Codec(), Header)
)

type codec struct{}

type decoder struct {
	r io.Reader
}

func (d decoder) Read(b []byte) (int, error) {
	encoded := make([]byte, hex.EncodedLen(len(b)))
	n, err := d.r.Read(encoded)
	if err != nil {
		return 0, err
	}
	return hex.Decode(b, encoded[:n])
}

func (d decoder) Decode(v interface{}) error {
	out, ok := v.(io.Writer)
	if !ok {
		return base.ErrExpectedWriter
	}

	_, err := io.Copy(out, d)
	return err
}

func (codec) Decoder(r io.Reader) mc.Decoder {
	return decoder{r}
}

type encoder struct {
	w io.Writer
}

func (e encoder) Write(b []byte) (int, error) {
	encoded := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(encoded, b)
	n, err := e.w.Write(encoded)
	return n / 2, err
}

func (e encoder) Encode(v interface{}) error {
	in, ok := v.(io.Reader)
	if !ok {
		return base.ErrExpectedReader
	}
	_, err := io.Copy(e, in)
	return err
}

func (codec) Encoder(w io.Writer) mc.Encoder {
	return encoder{w}
}

func Codec() mc.Codec {
	return codec{}
}

func Multicodec() mc.Multicodec {
	return multic
}
