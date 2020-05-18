package src

import (
	posmint_crypto "github.com/pokt-network/amino-js/go/lib/pokt-network/posmint/crypto"
	posmint_types "github.com/pokt-network/amino-js/go/lib/pokt-network/posmint/types"
	posmint_auth "github.com/pokt-network/amino-js/go/lib/pokt-network/posmint/x/auth"
)

func EncodePosmintTx(bz []byte, lengthPrefixed bool) (bz2 []byte, err error) {
	var o posmint_types.Tx
	err = codec.UnmarshalJSON(bz, &o)
	if err != nil {
		return nil, err
	}

	if lengthPrefixed {
		bz2, err = codec.MarshalBinaryLengthPrefixed(o)
	} else {
		bz2, err = codec.MarshalBinaryBare(o)
	}

	if err != nil {
		return nil, err
	}

	return
}

func EncodePosmintStdTx(bz []byte, lengthPrefixed bool) (bz2 []byte, err error) {
	var o posmint_auth.StdTx
	err = codec.UnmarshalJSON(bz, &o)
	if err != nil {
		return nil, err
	}

	if lengthPrefixed {
		bz2, err = codec.MarshalBinaryLengthPrefixed(o)
	} else {
		bz2, err = codec.MarshalBinaryBare(o)
	}

	if err != nil {
		return nil, err
	}

	return
}

func EncodePosmintStdSignDoc(bz []byte, lengthPrefixed bool) (bz2 []byte, err error) {
	var o posmint_auth.StdSignDoc
	err = codec.UnmarshalJSON(bz, &o)
	if err != nil {
		return nil, err
	}

	if lengthPrefixed {
		bz2, err = codec.MarshalBinaryLengthPrefixed(o)
	} else {
		bz2, err = codec.MarshalBinaryBare(o)
	}

	if err != nil {
		return nil, err
	}

	return
}

func EncodePosmintEd25519PublicKey(bz []byte, lengthPrefixed bool) (bz2 []byte, err error) {
	var o posmint_crypto.Ed25519PublicKey
	err = codec.UnmarshalJSON(bz, &o)
	if err != nil {
		return nil, err
	}

	if lengthPrefixed {
		bz2, err = codec.MarshalBinaryLengthPrefixed(o)
	} else {
		bz2, err = codec.MarshalBinaryBare(o)
	}

	if err != nil {
		return nil, err
	}

	return
}
