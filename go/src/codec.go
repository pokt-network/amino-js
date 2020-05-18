package src

import (
	"github.com/pokt-network/amino-js/go/lib"
	"github.com/tendermint/go-amino"
)

var codec *amino.Codec

func init() {
	codec = amino.NewCodec()
	lib.RegisterCodec(codec)
	codec.Seal()
}
