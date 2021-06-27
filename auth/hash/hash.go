package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"

	"golang.org/x/crypto/sha3"
)

type Type int8

const (
	SHA3_256 = Type(iota)
	SHA3_512
	SHA256
	SHA512
	SHA1
	MD5
)

func Hash(method Type, data []byte) []byte {
	var hashed []byte = nil
	switch method {
	case SHA1:
		r := sha1.Sum(data)
		hashed = r[:]
	case SHA3_256:
		r := sha3.Sum256(data)
		hashed = r[:]
	case SHA3_512:
		r := sha3.Sum512(data)
		hashed = r[:]
	case SHA256:
		r := sha256.Sum256(data)
		hashed = r[:]
	case SHA512:
		r := sha512.Sum512(data)
		hashed = r[:]
	case MD5:
		r := md5.Sum(data)
		hashed = r[:]
	}
	return hashed
}
