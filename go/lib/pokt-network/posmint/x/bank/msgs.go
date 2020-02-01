package bank

import (
	sdk "github.com/cosmos/amino-js/go/lib/pokt-network/posmint/types"
)

var _ sdk.Msg = MsgSend{}

// MsgSend - high level transaction of the coin module
type MsgSend struct {
	FromAddress sdk.Address `json:"from_address" yaml:"from_address"`
	ToAddress   sdk.Address `json:"to_address" yaml:"to_address"`
	Amount      sdk.Coins   `json:"amount" yaml:"amount"`
}
