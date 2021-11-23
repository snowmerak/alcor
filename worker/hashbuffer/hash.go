package hashbuffer

import (
	"alcor/db"
	"alcor/worker/hashbuffer/ringbuffer"
	"alcor/worker/rdb"
	"context"
	"crypto/sha512"
	"encoding/hex"
	"fmt"

	"github.com/snowmerak/jumper"
)

func init() {
	go Observe()
}

var buffer = ringbuffer.New(8)

func EnQueue(hash []byte) {
	buffer.EnQueue(hash)
}

type HashRingError struct {
	data string
}

func (h HashRingError) Error() string {
	return h.data
}

func NewRingError(data string) HashRingError {
	return HashRingError{data}
}

func Observe() {
	prev, err := rdb.GetLastBundle()
	if err != nil {
		prev = make([]byte, 32)
	}
	for {
		list := make([][]byte, 8)
		for i := 0; i < 8; i++ {
			list[i] = buffer.DeQueue()
		}
		bundle := new(db.Bundle)
		bundle.Prev = prev
		bundle.SubHashes = list
		sha := sha512.New()
		sha.Write(bundle.Prev)
		for _, v := range bundle.SubHashes {
			sha.Write(v)
		}
		bundle.Hash = sha.Sum(nil)
		hashed := sha512.Sum512(bundle.Hash)
		bundle.Hash = hashed[:]
		if err := db.InsertBundle(context.Background(), bundle); err != nil {
			jumper.Offer(NewRingError(err.Error()))
			fmt.Println(err)
			go func(list [][]byte) {
				for _, v := range list {
					EnQueue(v)
				}
			}(list)
			continue
		}
		fmt.Println("inserted bundle: ", hex.EncodeToString(bundle.Hash))
		fmt.Println("inserted bundle sub hashes: ")
		for _, v := range bundle.SubHashes {
			fmt.Println("- ", hex.EncodeToString(v))
		}
		rdb.SetLastBundle(bundle.Hash)
		prev = bundle.Hash
	}
}
