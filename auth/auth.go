package auth

import (
	"alcor/auth/hash"
	"crypto/elliptic"
)

var Basis = NewP384()

func NewP384() Ecdsa {
	return NewEcdsa(elliptic.P384(), hash.SHA512)
}
