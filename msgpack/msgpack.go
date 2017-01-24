package mc_msgpack

import (
	"io"
	"reflect"

	mc "github.com/multiformats/go-multicodec"
	gocodec "github.com/ugorji/go/codec"
)

var HeaderPath string
var Header []byte

// Handler options. See go/codec docs for their meanings.
var (
	// MsgpackHandle
	RawToString bool
	WriteExt    bool

	// BasicHandle
	TypeInfos *gocodec.TypeInfos

	// Encode options
	StructToArray       bool
	Canonical           bool
	CheckCircularRef    bool
	RecursiveEmptyCheck bool
	Raw                 bool
	AsSymbols           gocodec.AsSymbolFlag

	// Decode options
	MapType              reflect.Type
	SliceType            reflect.Type
	MaxInitLen           int
	ErrorIfNoField       bool
	ErrorIfNoArrayExpand bool
	SignedInteger        bool
	MapValueReset        bool
	InterfaceReset       bool
	InternString         bool
	PreferArrayOverSlice bool
)

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
	h.RawToString = RawToString
	h.WriteExt = WriteExt
	h.TypeInfos = TypeInfos
	h.StructToArray = StructToArray
	h.Canonical = Canonical
	h.CheckCircularRef = CheckCircularRef
	h.Raw = Raw
	h.AsSymbols = AsSymbols
	return &encoder{
		w:   w,
		mc:  c.mc,
		enc: gocodec.NewEncoder(w, h),
	}
}

func (c *codec) Decoder(r io.Reader) mc.Decoder {
	h := &gocodec.MsgpackHandle{}
	h.RawToString = RawToString
	h.WriteExt = WriteExt
	h.TypeInfos = TypeInfos
	h.MapType = MapType
	h.SliceType = SliceType
	h.MaxInitLen = MaxInitLen
	h.ErrorIfNoField = ErrorIfNoField
	h.SignedInteger = SignedInteger
	h.MapValueReset = MapValueReset
	h.InterfaceReset = InterfaceReset
	h.InternString = InternString
	h.PreferArrayOverSlice = PreferArrayOverSlice
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
