package wallet

import (
	"alcor/bufbox"
	"alcor/client"
	"encoding/hex"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

var Path = filepath.Join(".", "db", "wallet")

var store *bufbox.BufBox

func init() {
	store = new(bufbox.BufBoxWrapper).SetCompression(bufbox.BasisCompression).SetDirectory(Path).Unwrap()
}

func Put(a *client.Account) error {
	bs, err := proto.Marshal(a)
	if err != nil {
		return err
	}
	return store.Put(a.ID, bs)
}

func Delete(hash []byte) error {
	return store.Delete(hash)
}

func Get(hash []byte) (*client.Account, error) {
	bs, err := store.Get(hash)
	if err != nil {
		return nil, err
	}
	a := new(client.Account)
	if err := proto.Unmarshal(bs, a); err != nil {
		return nil, err
	}
	return a, nil
}

func List() ([]*client.Account, error) {
	list, err := os.ReadDir(Path)
	if err != nil {
		return nil, err
	}
	result := make([]*client.Account, 0, len(list))
	for _, v := range list {
		h, err := hex.DecodeString(v.Name())
		if err != nil {
			continue
		}
		a, err := Get(h)
		if err != nil {
			continue
		}
		result = append(result, a)
	}
	return result, nil
}
