package book

import (
	"alcor/bufbox"
	"alcor/client"
	"encoding/hex"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

var Path = filepath.Join(".", "db", "clients")

var store *bufbox.BufBox

func init() {
	store = bufbox.NewWrapper().SetCompression(bufbox.BasisCompression).SetDirectory(Path).Unwrap()
}

func Put(c *client.Client) error {
	bs, err := proto.Marshal(c)
	if err != nil {
		return err
	}
	return store.Put(c.ID, bs)
}

func Get(hash []byte) (*client.Client, error) {
	bs, err := store.Get(hash)
	if err != nil {
		return nil, err
	}
	c := new(client.Client)
	if err := proto.Unmarshal(bs, c); err != nil {
		return nil, err
	}
	return c, nil
}

func List() ([]*client.Client, error) {
	list, err := os.ReadDir(Path)
	if err != nil {
		return nil, err
	}
	result := make([]*client.Client, 0, len(list))
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
