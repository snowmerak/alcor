package transaction

import "crypto/sha512"

func MakeHash(t *Transaction) []byte {
	bs := make([]byte, len(t.From)+len(t.Key)+len(t.Message)+len(t.Timestamp))
	bs = append(bs, t.From...)
	bs = append(bs, t.Key...)
	bs = append(bs, t.Message...)
	bs = append(bs, t.Timestamp...)
	h := sha512.Sum512(bs)
	return h[:]
}
