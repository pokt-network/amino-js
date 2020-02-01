package crypto

var _ PublicKey = Ed25519PublicKey{}

type Ed25519PublicKey [32]byte

// import (
// 	"github.com/cosmos/amino-js/go/lib/tendermint/tendermint/crypto/ed25519"
// )

// type (
// 	Ed25519PublicKey ed25519.PubKeyEd25519
// )

// var (
// 	_ PublicKey = Ed25519PublicKey{}
// )
