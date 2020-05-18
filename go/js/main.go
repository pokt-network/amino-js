package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	exports := js.Module.Get("exports")

	// @formatter:off
	// Basic encoding
	exports.Set("encodeByte", EncodeByte)
	exports.Set("encodeByteSlice", EncodeByteSlice)
	exports.Set("encodeInt8", EncodeInt8)
	exports.Set("encodeInt16", EncodeInt16)
	exports.Set("encodeInt32", EncodeInt32)
	exports.Set("encodeInt64", EncodeInt64)
	exports.Set("encodeVarint", EncodeVarint)
	exports.Set("encodeUint8", EncodeUint8)
	exports.Set("encodeUint16", EncodeUint16)
	exports.Set("encodeUint32", EncodeUint32)
	exports.Set("encodeUint64", EncodeUint64)
	exports.Set("encodeUvarint", EncodeUvarint)
	exports.Set("encodeFloat32", EncodeFloat32)
	exports.Set("encodeFloat64", EncodeFloat64)
	exports.Set("encodeBool", EncodeBool)
	exports.Set("encodeString", EncodeString)
	exports.Set("encodeTime", EncodeTime)

	// Meta
	exports.Set("decodeDisambPrefixBytes", DecodeDisambPrefixBytes)
	exports.Set("nameToDisfix", NameToDisfix)
	exports.Set("byteSliceSize", ByteSliceSize)
	exports.Set("uvarintSize", UvarintSize)
	exports.Set("varintSize", VarintSize)

	// Typed encoding
	exports.Set("encodePosmintTx", EncodePosmintTx)
	exports.Set("encodePosmintStdSignDoc", EncodePosmintStdSignDoc)
	exports.Set("encodePosmintStdTx", EncodePosmintStdTx)
	exports.Set("encodePosmintEd25519PublicKey", EncodePosmintEd25519PublicKey)
	// @formatter:on
}
