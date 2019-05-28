package types

import (
	"bytes"
	"sync"

	"github.com/jordansexton/go-amino-js/go/types/tendermint/tendermint/crypto/merkle"
	cmn "github.com/jordansexton/go-amino-js/go/types/tendermint/tendermint/libs/common"
)

type Part struct {
	Index int                `json:"index"`
	Bytes cmn.HexBytes       `json:"bytes"`
	Proof merkle.SimpleProof `json:"proof"`
}

type PartSetHeader struct {
	Total int          `json:"total"`
	Hash  cmn.HexBytes `json:"hash"`
}

type PartSet struct {
	total int
	hash  []byte

	mtx           sync.Mutex
	parts         []*Part
	partsBitArray *cmn.BitArray
	count         int
}

type PartSetReader struct {
	i      int
	parts  []*Part
	reader *bytes.Reader
}