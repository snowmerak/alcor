package tangle

import (
	"alcor/bufbox"
	"alcor/transaction"
	"math/rand"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

var Path = filepath.Join(".", "db", "tangle")

var store *bufbox.BufBox

func init() {
	store = new(bufbox.BufBoxWrapper).SetCompression(bufbox.BasisCompression).SetDirectory(Path).Unwrap()
}

func Put(t *transaction.Transaction) error {
	bs, err := proto.Marshal(t)
	if err != nil {
		return err
	}
	return store.Put(t.Hash, bs)
}

func Get(hash []byte) (*transaction.Transaction, error) {
	bs, err := store.Get(hash)
	if err != nil {
		return nil, err
	}
	t := new(transaction.Transaction)
	if err := proto.Unmarshal(bs, t); err != nil {
		return nil, err
	}
	return t, nil
}

func GetRandoms() (*transaction.Transaction, *transaction.Transaction, error) {
	list, err := store.List()
	if err != nil {
		return nil, nil, err
	}
	i, j := rand.Intn(len(list)), rand.Intn(len(list))
	if len(list) == 1 {
		t1, err := Get(list[i])
		if err != nil {
			return nil, nil, err
		}
		return t1, nil, err
	}
	for i != j {
		j = rand.Intn(len(list))
	}
	t1, err := Get(list[i])
	if err != nil {
		return nil, nil, err
	}
	t2, err := Get(list[j])
	if err != nil {
		return nil, nil, err
	}
	return t1, t2, err
}
