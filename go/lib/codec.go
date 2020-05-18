package lib

import (
	posmint_crypto "github.com/pokt-network/amino-js/go/lib/pokt-network/posmint/crypto"
	posmint_types "github.com/pokt-network/amino-js/go/lib/pokt-network/posmint/types"
	posmint_auth "github.com/pokt-network/amino-js/go/lib/pokt-network/posmint/x/auth"
	posmint_bank "github.com/pokt-network/amino-js/go/lib/pokt-network/posmint/x/bank"
	tm_crypto "github.com/pokt-network/amino-js/go/lib/tendermint/tendermint/crypto"

	pocket_core_apps "github.com/pokt-network/amino-js/go/lib/pokt-network/pocket-core/x/apps/types"
	pocket_core_nodes "github.com/pokt-network/amino-js/go/lib/pokt-network/pocket-core/x/nodes/types"

	amino "github.com/tendermint/go-amino"
)

func RegisterCodec(codec *amino.Codec) {
	// @formatter:off

	// tendermint/tendermint/crypto
	codec.RegisterInterface((*tm_crypto.PubKey)(nil), nil)

	// pokt-network/posmint/crypto
	codec.RegisterInterface((*posmint_crypto.PublicKey)(nil), nil)
	codec.RegisterConcrete(posmint_crypto.Ed25519PublicKey{}, PosmintEd25519PublicKey, nil)

	// pokt-network/posmint/types/codec.go

	codec.RegisterInterface((*posmint_types.Msg)(nil), nil)
	codec.RegisterInterface((*posmint_types.Tx)(nil), nil)

	// pokt-network/posmint/x/auth

	codec.RegisterConcrete(posmint_auth.StdTx{}, PosmintStdTx, nil)
	codec.RegisterConcrete(posmint_auth.StdSignDoc{}, PosmintStdSignDoc, nil)

	// pokt-network/posmint/x/bank

	codec.RegisterConcrete(posmint_bank.MsgSend{}, PosmintMsgSend, nil)

	// pokt-network/pocket-core/x/apps

	codec.RegisterConcrete(pocket_core_apps.MsgAppStake{}, PocketCoreMsgAppStake, nil)
	codec.RegisterConcrete(pocket_core_apps.MsgBeginAppUnstake{}, PocketCoreMsgBeginAppUnstake, nil)
	codec.RegisterConcrete(pocket_core_apps.MsgAppUnjail{}, PocketCoreMsgAppUnjail, nil)

	// pokt-network/pocket-core/x/nodes

	codec.RegisterConcrete(pocket_core_nodes.MsgStake{}, PocketCoreMsgStake, nil)
	codec.RegisterConcrete(pocket_core_nodes.MsgBeginUnstake{}, PocketCoreMsgBeginUnstake, nil)
	codec.RegisterConcrete(pocket_core_nodes.MsgUnjail{}, PocketCoreMsgUnjail, nil)

	// @formatter:on
}
