package crypto

type Ed25519PublicKey [32]byte

var _ PublicKey = Ed25519PublicKey{}
