package hashbuffer

import (
	"alcor/db"
	"alcor/worker/hashbuffer/ringbuffer"
	"context"
	"crypto/sha512"

	"github.com/snowmerak/jumper"
)

func init() {
	go Observe()
}

var buffer = ringbuffer.New(8)

func EnQueue(hash []byte) {
	buffer.EnQueue(hash)
}

func DeQueue(pool [][]byte) {
	for i := 0; i < len(pool); i++ {
		pool[i] = buffer.DeQueue()
	}
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
	prev := []byte{}
	for {
		list := make([][]byte, 8)
		DeQueue(list)
		bundle := new(db.Bundle)
		bundle.Prev = prev
		bundle.SubHashes = list
		sha := sha512.New()
		sha.Write(bundle.Prev)
		for _, v := range bundle.SubHashes {
			sha.Write(v)
		}
		bundle.Hash = sha.Sum(nil)
		if err := db.InsertBundle(context.Background(), bundle); err != nil {
			jumper.Offer(NewRingError(err.Error()))
			go func(list [][]byte) {
				for _, v := range list {
					EnQueue(v)
				}
			}(list)
			continue
		}
		prev = bundle.Prev
	}
}
