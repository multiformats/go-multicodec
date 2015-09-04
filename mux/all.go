package muxcodec

import (
	mc "github.com/jbenet/go-multicodec"
	cbor "github.com/jbenet/go-multicodec/cbor"
	json "github.com/jbenet/go-multicodec/json"
)

func StandardMux() *Multicodec {
	return MuxMulticodec([]mc.Multicodec{
		cbor.Multicodec(),
		json.Multicodec(false),
		json.Multicodec(true),
	}, SelectFirst)
}
