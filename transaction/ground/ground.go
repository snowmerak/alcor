package ground

import (
	"alcor/bufbox"
	"alcor/transaction"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

var Path = filepath.Join(".", "db", "transactions")

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
