package types

import (
	sdk "github.com/jordansexton/go-amino-js/go/types/cosmos/cosmos-sdk/types"
)

type Vote struct {
	Voter      sdk.AccAddress `json:"voter"`       //  address of the voter
	ProposalID uint64         `json:"proposal_id"` //  proposalID of the proposal
	Option     VoteOption     `json:"option"`      //  option from OptionSet chosen by the voter
}

type Votes []Vote

type VoteOption byte
