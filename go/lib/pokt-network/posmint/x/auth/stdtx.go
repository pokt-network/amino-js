package auth

import (
	"encoding/json"

	sdk "github.com/pokt-network/amino-js/go/lib/pokt-network/posmint/types"
	"github.com/pokt-network/amino-js/go/lib/tendermint/tendermint/crypto"
)

var _ sdk.Tx = (*StdTx)(nil)

type StdTx struct {
	Msgs       []sdk.Msg      `json:"msg"`
	Fee        sdk.Coins      `json:"fee"`
	Signatures []StdSignature `json:"signatures"`
	Memo       string         `json:"memo"`
	Entropy    int64          `json:"entropy"`
}

type StdSignDoc struct {
	ChainID       string            `json:"chain_id"`
	Fee           json.RawMessage   `json:"fee"`
	Memo          string            `json:"memo"`
	Msgs          []json.RawMessage `json:"msgs"`
	Entropy       int64             `json:"entropy"`
}

type StdSignature struct {
	crypto.PubKey `json:"pub_key"` // optional
	Signature     []byte           `json:"signature"`
}
