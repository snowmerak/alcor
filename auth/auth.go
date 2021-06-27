package auth

import (
	"crypto/ecdsa"
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
