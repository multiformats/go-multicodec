package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	mc "github.com/jbenet/go-multicodec"
	cbor "github.com/jbenet/go-multicodec/cbor"
	json "github.com/jbenet/go-multicodec/json"
	mux "github.com/jbenet/go-multicodec/mux"
)

// flags
var Flags FlagType

type FlagType struct {
	Recode string
	Filter string
	Wrap   string

	Msgio   bool
	MCWrap  bool
	Headers bool
	Paths   bool

	P2H bool
	H2P bool
}

var usage = `
multicodec - tool to inspect and manipulate mixed codec streams

Usage
  cat rawjson | multicodec --wrap /json/msgio >mcjson
  cat rawcbor | multicodec --wrap /cbor >mccbor

  cat mixed | multicodec -r /json/msgio >all_in_json
  cat mixed | multicodec -f /json/msgio >json_ones_only

  cat mixed | multicodec --headers >all_headers
  cat mixed | multicodec --paths >all_paths

  cat paths   | multicodec --p2h >headers
  cat headers | multicodec --h2p >paths

Options
  -f, --filter    filter items of given codec
  -r, --recode    recode items to given codec
  -w, --wrap      wrap raw data with header

  --mcwrap        item headers wrapped with /multicodec
  --msgio         wrap all subcodecs with /msgio

  --headers       output only items' headers
  --paths         output only items' header paths

  --h2p           convert headers to line-delimited paths
  --p2h           convert line-delimited paths to headers
`

func init() {
	flag.StringVar(&Flags.Recode, "r", "", "recode items to given codec")
	flag.StringVar(&Flags.Recode, "recode", "", "recode items to given codec")
	flag.StringVar(&Flags.Recode, "f", "", "filter items of given codec")
	flag.StringVar(&Flags.Recode, "filter", "", "filter items of given codec")
	flag.StringVar(&Flags.Wrap, "w", "", "wrap raw data with header")
	flag.StringVar(&Flags.Wrap, "wrap", "", "wrap raw data with header")
	flag.BoolVar(&Flags.MCWrap, "mcwrap", false, "items headers wrapped with /multicodec")
	flag.BoolVar(&Flags.Msgio, "msgio", false, "wrap all subcodecs with /msgio")
	flag.BoolVar(&Flags.Headers, "headers", false, "output only the headers")
	flag.BoolVar(&Flags.Paths, "paths", false, "output only the header paths")
	flag.BoolVar(&Flags.P2H, "p2h", false, "paths to headers")
	flag.BoolVar(&Flags.H2P, "h2p", false, "headers to paths")
	flag.Usage = func() {
		fmt.Println(strings.TrimSpace(usage))
		os.Exit(0)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(-1)
	}
}

func run() error {
	flag.Parse()

	w := os.Stdout
	r := os.Stdin

	switch {
	case Flags.Headers:
		return headers(w, r)
	case Flags.Paths:
		return paths(w, r)
	case Flags.Wrap != "":
		return wrap(w, r, Flags.Wrap)
	case Flags.Filter != "":
		return filter(w, r, Flags.Filter)
	case Flags.Recode != "":
		return recode(w, r, Flags.Recode)
	case Flags.H2P:
		return h2p(w, r)
	case Flags.P2H:
		return p2h(w, r)
	default:
		flag.Usage()
		return nil
	}
}

func h2p(w io.Writer, r io.Reader) error {
	for {
		hdr, err := mc.ReadHeader(r)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		p := string(mc.HeaderPath(hdr))

		_, err = fmt.Fprintln(w, p)
		if err != nil {
			return err
		}
	}
}

func p2h(w io.Writer, r io.Reader) error {
	br := bufio.NewReader(r)

	for {
		p, err := br.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		p = p[:len(p)-1] // remove \n

		hdr := mc.Header(p)

		_, err = w.Write(hdr)
		if err != nil {
			return err
		}
	}
}

func headers(w io.Writer, r io.Reader) error {
	return decode(r, func(codec *mux.Multicodec, v interface{}) error {
		_, err := w.Write(codec.Last.Header())
		return err
	})
}

func paths(w io.Writer, r io.Reader) error {
	return decode(r, func(codec *mux.Multicodec, v interface{}) error {
		p := mc.HeaderPath(codec.Last.Header())
		_, err := fmt.Fprintln(w, string(p))
		return err
	})
}

func wrap(w io.Writer, r io.Reader, path string) error {
	mcc := codecWithPath(path)
	if mcc == nil {
		return fmt.Errorf("unknown codec %s", path)
	}
	hdrs := string(mcc.Header())

	wrapRT := func(cc mc.Codec, mcc mc.Multicodec) error {
		var v interface{}
		if err := cc.Decoder(r).Decode(&v); err != nil {
			return err
		}
		return mcc.Encoder(w).Encode(&v)
	}

	switch hdrs {
	case string(json.HeaderMsgio):
		return wrapRT(json.Codec(true), mcc)
	case string(json.Header):
		return wrapRT(json.Codec(false), mcc)
	case string(cbor.Header):
		return wrapRT(cbor.Codec(), mcc)
	}

	return fmt.Errorf("wrap unsupported for codec %s", hdrs)
}

func filter(w io.Writer, r io.Reader, path string) error {
	hdr := mc.Header([]byte(path))

	// as we decode, write everything to a buffer.
	buf := bytes.NewBuffer(nil)
	r = io.TeeReader(r, buf)

	return decode(r, func(codec *mux.Multicodec, value interface{}) error {
		defer buf.Reset()

		// c.Last is the last multicodec encoded or decoded.
		// skip headers which do not match.
		if !bytes.Equal(hdr, codec.Last.Header()) {
			return nil
		}

		_, err := io.Copy(w, buf)
		if err != nil && err != io.EOF {
			return err
		}
		return nil
	})
}

func recode(w io.Writer, r io.Reader, path string) error {
	codec := codecWithPath(path)
	if codec == nil {
		return fmt.Errorf("unknown codec %s", path)
	}
	enc := codec.Encoder(w)

	return decode(r, func(codec *mux.Multicodec, v interface{}) error {
		return enc.Encode(v)
	})
}

func decode(r io.Reader, next func(m *mux.Multicodec, v interface{}) error) error {
	c := codec()
	dec := c.Decoder(r)

	for {
		var v interface{}
		if err := dec.Decode(&v); err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		if err := next(c, v); err != nil {
			return err
		}
	}

	return nil
}

func codecWithPath(path string) mc.Multicodec {
	hdr := mc.Header([]byte(path))
	codec := mux.CodecWithHeader(hdr, mux.StandardMux().Codecs)
	return codec
}

func codec() *mux.Multicodec {
	m := mux.StandardMux()
	m.Wrap = Flags.MCWrap
	return m
}
