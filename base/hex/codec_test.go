package bin

import (
	"bytes"
	"testing"

	mc "github.com/multiformats/go-multicodec"
)

func TestHexDecoding(t *testing.T) {
	bufIn := bytes.Buffer{}
	bufIn.Write(Multicodec().Header())
	dataIn := []byte{255, 255}
	bufIn.Write([]byte("ffff"))

	dataOut := &bytes.Buffer{}
	err := Multicodec().Decoder(&bufIn).Decode(dataOut)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(dataIn, dataOut.Bytes()) {
		t.Fatalf("dataOut(%v) is not eqal to dataIn(%v)", dataOut, dataIn)
	}
}

func TestHexEncoding(t *testing.T) {
	buf := bytes.Buffer{}
	data := []byte("ffff")

	err := Multicodec().Encoder(&buf).Encode(bytes.NewReader([]byte{255, 255}))
	if err != nil {
		t.Fatal(err)
	}

	err = mc.ConsumeHeader(&buf, Multicodec().Header())
	if err != nil {
		t.Fatal(err)
	}

	dataOut := make([]byte, len(data))
	buf.Read(dataOut)

	if !bytes.Equal(data, dataOut) {
		t.Fatalf("dataOut(%v) is not eqal to data(%v)", dataOut, data)
	}
}
