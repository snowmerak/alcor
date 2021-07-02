package auth

import (
	"alcor/auth/hash"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"errors"
	"math/big"

	"google.golang.org/protobuf/proto"
)

type Ecdsa struct {
	curve      elliptic.Curve
	hashMethod hash.Type
}

func NewEcdsa(curve elliptic.Curve, hashMethod hash.Type) Ecdsa {
	return Ecdsa{curve: curve}
}

func (e Ecdsa) GeneratePrivateKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(e.curve, rand.Reader)
}

func (e Ecdsa) GeneratePublicKey(pr *ecdsa.PrivateKey) (*ecdsa.PublicKey, error) {
	if pr == nil {
		return nil, errors.New("privatekey is nil")
	}
	return &pr.PublicKey, nil
}

func (e Ecdsa) Sign(pr *ecdsa.PrivateKey, data []byte) (*big.Int, *big.Int, error) {
	hashed := hash.Hash(e.hashMethod, data)
	return ecdsa.Sign(rand.Reader, pr, hashed)
}

func (e Ecdsa) Verify(pb *ecdsa.PublicKey, data []byte, r *big.Int, s *big.Int) bool {
	hashed := hash.Hash(e.hashMethod, data)
	return ecdsa.Verify(pb, hashed, r, s)
}

func (e Ecdsa) SerializePrivateKey(pr *ecdsa.PrivateKey) ([]byte, error) {
	return x509.MarshalECPrivateKey(pr)
}

func (e Ecdsa) DeserializePrivateKey(bs []byte) (*ecdsa.PrivateKey, error) {
	return x509.ParseECPrivateKey(bs)
}

func (e Ecdsa) SerializePublicKey(pb *ecdsa.PublicKey) ([]byte, error) {
	return x509.MarshalPKIXPublicKey(pb)
}

func (e Ecdsa) DeserializePublicKey(bs []byte) (*ecdsa.PublicKey, error) {
	pb, err := x509.ParsePKIXPublicKey(bs)
	if err != nil {
		return nil, err
	}
	rs, ok := pb.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("converted value is not ecdsa public key pointer")
	}
	return rs, nil
}

func (e Ecdsa) SerializeSignature(r *big.Int, s *big.Int) ([]byte, error) {
	sig := &Signature{
		R: r.Text(16),
		S: s.Text(16),
	}
	return proto.Marshal(sig)
}

func (e Ecdsa) DeserializeSignature(data []byte) (*big.Int, *big.Int, error) {
	sig := new(Signature)
	if err := proto.Unmarshal(data, sig); err != nil {
		return nil, nil, err
	}
	r, ok := new(big.Int).SetString(sig.R, 16)
	if !ok {
		return nil, nil, errors.New("bad value of R")
	}
	s, ok := new(big.Int).SetString(sig.S, 16)
	if !ok {
		return nil, nil, errors.New("bad value of S")
	}
	return r, s, nil
}
