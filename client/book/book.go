package book

import (
	"alcor/bufbox"
	"alcor/client"
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
