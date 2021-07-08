package transaction

import "crypto/sha512"

func MakePair(hash1 []byte, hash2 []byte) *Pair {
	p := new(Pair)
	p.Insides = append(p.Insides, hash1, hash2)
	bs := []byte(nil)
	for _, h := range p.Insides {
		bs = append(bs, h...)
		th := sha512.Sum512(bs)
		bs = th[:]
	}
	p.Hash = bs
	return p
}
