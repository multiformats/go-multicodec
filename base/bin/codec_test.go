package bin

import (
	"bytes"
	"testing"

	mc "github.com/multiformats/go-multicodec"
)

func TestBinaryDecoding(t *testing.T) {
	buf := bytes.Buffer{}
	buf.Write(Header)
	data := []byte("Multicodec")
	buf.Write(data)

	dataOut := bytes.Buffer{}
	err := Multicodec().Decoder(&buf).Decode(&dataOut)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, dataOut.Bytes()) {
		t.Fatalf("dataOut(%v) is not eqal to data(%v)", dataOut, data)
	}
}

func TestBinaryEncoding(t *testing.T) {
	buf := bytes.Buffer{}
	data := []byte("Is Awesome")

	err := Multicodec().Encoder(&buf).Encode(bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	err = mc.ConsumeHeader(&buf, Header)
	if err != nil {
		t.Fatal(err)
	}

	dataOut := make([]byte, len(data))
	buf.Read(dataOut)

	if !bytes.Equal(data, dataOut) {
		t.Fatalf("dataOut(%v) is not eqal to data(%v)", dataOut, data)
	}
}
