package auth

import (
	"alcor/model/auth"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"math/big"

	"google.golang.org/protobuf/proto"
)

func SerializePrivateKey(pr *ecdsa.PrivateKey) ([]byte, error) {
	prpb := new(auth.PrivateKey)
	return proto.Marshal(prpb)
}

func DeserializePrivateKey(bs []byte) (*ecdsa.PrivateKey, error) {
	pr, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	prpb := new(auth.PrivateKey)
	if err := proto.Unmarshal(bs, prpb); err != nil {
		return nil, err
	}
	if _, err := pr.X.SetString(prpb.X, 16); !err {
		return nil, errors.New("cannot convert")
	}
	if _, err := pr.Y.SetString(prpb.Y, 16); !err {
		return nil, errors.New("cannot convert")
	}
	return pr, nil
}

func SerializePublicKey(pb *ecdsa.PublicKey) ([]byte, error) {
	pbpb := new(auth.PublicKey)
	pbpb.X = pb.X.Text(16)
	pbpb.Y = pb.Y.Text(16)
	return proto.Marshal(pbpb)
}

func DeserializePublicKey(bs []byte) (*ecdsa.PublicKey, error) {
	pr, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	pb := pr.PublicKey
	pbpb := new(auth.PublicKey)
	if err := proto.Unmarshal(bs, pbpb); err != nil {
		return nil, err
	}
	if _, err := pb.X.SetString(pbpb.X, 16); !err {
		return nil, errors.New("cannot convert")
	}
	if _, err := pb.Y.SetString(pbpb.Y, 16); !err {
		return nil, errors.New("cannot convert")
	}
	return &pb, nil
}

func SerializeSignature(r *big.Int, s *big.Int) ([]byte, error) {
	sig := &auth.Signature{
		R: r.Text(16),
		S: s.Text(16),
	}
	return proto.Marshal(sig)
}

func DeserializeSignature(data []byte) (*big.Int, *big.Int, error) {
	sig := new(auth.Signature)
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
