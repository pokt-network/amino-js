package auth

import (
	"encoding/json"

	sdk "github.com/cosmos/amino-js/go/lib/pokt-network/posmint/types"
	"github.com/cosmos/amino-js/go/lib/tendermint/tendermint/crypto"
)

var _ sdk.Tx = (*StdTx)(nil)

type StdTx struct {
	Msgs       []sdk.Msg      `json:"msg"`
	Fee        sdk.Coins      `json:"fee"`
	Signatures []StdSignature `json:"signatures"`
	Memo       string         `json:"memo"`
}

type StdSignDoc struct {
	AccountNumber uint64            `json:"account_number"`
	ChainID       string            `json:"chain_id"`
	Fee           json.RawMessage   `json:"fee"`
	Memo          string            `json:"memo"`
	Msgs          []json.RawMessage `json:"msgs"`
	Sequence      uint64            `json:"sequence"`
}

type StdSignature struct {
	crypto.PubKey `json:"pub_key"` // optional
	Signature     []byte           `json:"signature"`
}
