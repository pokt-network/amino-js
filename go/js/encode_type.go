package main

import (
	"github.com/pokt-network/amino-js/go/src"
)

func EncodePosmintTx(bz []byte, lengthPrefixed bool) (bz2 []byte) {
	bz2, err := src.EncodePosmintTx(bz, lengthPrefixed)
	if err != nil {
		panic(err)
	}
	return
}

func EncodePosmintStdTx(bz []byte, lengthPrefixed bool) (bz2 []byte) {
	bz2, err := src.EncodePosmintStdTx(bz, lengthPrefixed)
	if err != nil {
		panic(err)
	}
	return
}

func EncodePosmintStdSignDoc(bz []byte, lengthPrefixed bool) (bz2 []byte) {
	bz2, err := src.EncodePosmintStdSignDoc(bz, lengthPrefixed)
	if err != nil {
		panic(err)
	}
	return
}

func EncodePosmintEd25519PublicKey(bz []byte, lengthPrefixed bool) (bz2 []byte) {
	bz2, err := src.EncodePosmintEd25519PublicKey(bz, lengthPrefixed)
	if err != nil {
		panic(err)
	}
	return
}
