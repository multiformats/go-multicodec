package bin

import (
	"io"

	mc "github.com/multiformats/go-multicodec"
	base "github.com/multiformats/go-multicodec/base"
)

var (
	HeaderPath = "/bin/"
	Header     = mc.Header([]byte(HeaderPath))
	multic     = mc.NewMulticodecFromCodec(Codec(), Header)
)

type codec struct{}

func (codec) Header() []byte {
	return Header
}

type decoder struct {
	r io.Reader
}

func (d decoder) Decode(v interface{}) error {
	out, ok := v.(io.Writer)
	if !ok {
		return base.ErrExpectedWriter
	}

	_, err := io.Copy(out, d.r)
	return err
}

func (codec) Decoder(r io.Reader) mc.Decoder {
	return decoder{r}
}

type encoder struct {
	w io.Writer
}

func (e encoder) Encode(v interface{}) error {
	in, ok := v.(io.Reader)
	if !ok {
		return base.ErrExpectedReader
	}

	_, err := io.Copy(e.w, in)
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
