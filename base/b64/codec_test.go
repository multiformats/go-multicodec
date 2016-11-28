package b64

import (
	"bytes"
	"testing"

	mc "github.com/multiformats/go-multicodec"
)

func TestB64Decoding(t *testing.T) {
	bufIn := bytes.Buffer{}
	bufIn.Write(Multicodec().Header())
	dataIn := []byte{255, 255}
	bufIn.Write([]byte("//8="))

	dataW := new(bytes.Buffer)
	err := Multicodec().Decoder(&bufIn).Decode(dataW)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(dataIn, dataW.Bytes()) {
		t.Fatalf("dataOut(%v) is not eqal to dataIn(%v)", dataW.Bytes(), dataIn)
	}
}

func TestB64Encoding(t *testing.T) {
	buf := bytes.Buffer{}
	data := []byte("//8=")

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
