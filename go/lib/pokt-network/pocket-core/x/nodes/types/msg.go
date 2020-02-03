package types

import (
	"github.com/cosmos/amino-js/go/lib/pokt-network/posmint/crypto"
	sdk "github.com/cosmos/amino-js/go/lib/pokt-network/posmint/types"
)

var (
	_ sdk.Msg = &MsgStake{}
	_ sdk.Msg = &MsgBeginUnstake{}
)

// MsgStake - struct for staking transactions
type MsgStake struct {
	PublicKey  crypto.PublicKey `json:"public_key" yaml:"public_key"`
	Chains     []string         `json:"chains" yaml:"chains"`
	Value      sdk.Int          `json:"value" yaml:"value"`
	ServiceURL string           `json:"service_url" yaml:"service_url"`
}

// MsgBeginUnstake - struct for unstaking transaciton
type MsgBeginUnstake struct {
	Address sdk.Address `json:"validator_address" yaml:"validator_address"`
}

// MsgUnjail - struct for unjailing jailed validator
type MsgUnjail struct {
	ValidatorAddr sdk.Address `json:"address" yaml:"address"` // address of the validator operator
}
