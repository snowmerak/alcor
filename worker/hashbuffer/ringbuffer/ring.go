package ringbuffer

import (
	"runtime"
	"sync/atomic"
)

/*
This code is created my lemon-mint.
I edited the type unsafe.Pointer to []byte.
*/

//RingBuffer ...
type RingBuffer struct {
	buf     [][]byte
	size    int
	r, w    int
	counter int64
	TLock
}

func New(size int) *RingBuffer {
	r := new(RingBuffer)
	r.buf = make([][]byte, size+1)
	r.size = size + 1
	return r
}

func (b *RingBuffer) EnQueue(x []byte) {
	for {
		ctr := b.counter
		if ctr+1 >= int64(b.size) {
			runtime.Gosched()
			continue
		}
		if atomic.CompareAndSwapInt64(&b.counter, ctr, ctr+1) {
			break
		}
	}
	b.Lock()
	b.buf[b.w] = x
	b.w++
	if b.w >= b.size {
		b.w = 0
	}
	b.Unlock()
}

func (b *RingBuffer) DeQueue() []byte {
	for {
		ctr := b.counter
		if ctr <= 0 {
			runtime.Gosched()
			continue
		}
		if atomic.CompareAndSwapInt64(&b.counter, ctr, ctr-1) {
			break
		}
	}
	b.Lock()
	val := b.buf[b.r]
	b.r++
	if b.r >= b.size {
		b.r = 0
	}
	b.Unlock()
	return val
}
