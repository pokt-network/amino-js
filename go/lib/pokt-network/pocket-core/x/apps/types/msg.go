package types

import (
	"github.com/cosmos/amino-js/go/lib/pokt-network/posmint/crypto"
	sdk "github.com/cosmos/amino-js/go/lib/pokt-network/posmint/types"
)

var (
	_ sdk.Msg = &MsgAppStake{}
	_ sdk.Msg = &MsgBeginAppUnstake{}
)

// MsgAppStake - struct for staking transactions
type MsgAppStake struct {
	PubKey crypto.PublicKey `json:"pubkey" yaml:"pubkey"`
	Chains []string         `json:"chains" yaml:"chains"`
	Value  sdk.Int          `json:"value" yaml:"value"`
}

// MsgBeginAppUnstake - struct for unstaking transaciton
type MsgBeginAppUnstake struct {
	Address sdk.Address `json:"application_address" yaml:"application_address"`
}
