package auth

import (
	"alcor/auth/hash"
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
)

type Auth interface {
	GeneratePrivateKey() (*ecdsa.PrivateKey, error)
	GeneratePublicKey(*ecdsa.PrivateKey) (*ecdsa.PublicKey, error)
	Sign(*ecdsa.PrivateKey, []byte) (*big.Int, *big.Int, error)
	Verify(*ecdsa.PublicKey, []byte, *big.Int, *big.Int) bool
	SerializePrivateKey(*ecdsa.PrivateKey) ([]byte, error)
	DeserializePrivateKey([]byte) (*ecdsa.PrivateKey, error)
	SerializePublicKey(*ecdsa.PublicKey) ([]byte, error)
	DeserializePublicKey([]byte) (*ecdsa.PublicKey, error)
	SerializeSignature(*big.Int, *big.Int) ([]byte, error)
	DeserializeSignature([]byte) (*big.Int, *big.Int, error)
}

var Basis = NewP384()

func NewP384() Auth {
	return NewEcdsa(elliptic.P384(), hash.SHA512)
}
