package bank

import (
	sdk "github.com/pokt-network/amino-js/go/lib/pokt-network/posmint/types"
)

var _ sdk.Msg = MsgSend{}

// MsgSend - high level transaction of the coin module
type MsgSend struct {
	FromAddress sdk.Address
	ToAddress   sdk.Address
	Amount      sdk.Int
}
