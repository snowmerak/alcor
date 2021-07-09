package pyramid

import (
	"alcor/bufbox"
	"alcor/transaction"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

var Path = filepath.Join(".", "db", "pairs")

var store *bufbox.BufBox

func init() {
	store = bufbox.NewWrapper().SetCompression(bufbox.BasisCompression).SetDirectory(Path).Unwrap()
}

func Put(p *transaction.Pair) error {
	bs, err := proto.Marshal(p)
	if err != nil {
		return err
	}
	return store.Put(p.Hash, bs)
}

func Get(hash []byte) (*transaction.Pair, error) {
	bs, err := store.Get(hash)
	if err != nil {
		return nil, err
	}
	p := new(transaction.Pair)
	if err := proto.Unmarshal(bs, p); err != nil {
		return nil, err
	}
	return p, nil
}
