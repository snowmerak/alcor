package wsstation_test

import (
	"sync"
	"testing"
)

type Signal struct {
	in int8
}

func BenchmarkValue(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan Signal)
	go func() {
		for i := 0; i < 1000; i++ {
			s := Signal{}
			ch <- s
		}
		close(ch)
		wg.Done()
	}()
	go func() {
		for v := range ch {
			_ = v
		}
		wg.Done()
	}()
	wg.Wait()
}

func BenchmarkValueBuf8(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan Signal, 8)
	go func() {
		for i := 0; i < 1000; i++ {
			s := Signal{}
			ch <- s
		}
		close(ch)
		wg.Done()
	}()
	go func() {
		for v := range ch {
			_ = v
		}
		wg.Done()
	}()
	wg.Wait()
}

func BenchmarkPointer(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan *Signal)
	go func() {
		for i := 0; i < 1000; i++ {
			s := &Signal{}
			ch <- s
		}
		close(ch)
		wg.Done()
	}()
	go func() {
		for v := range ch {
			_ = *v
		}
		wg.Done()
	}()
	wg.Wait()
}

func BenchmarkPointerBuf8(b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan *Signal, 8)
	go func() {
		for i := 0; i < 1000; i++ {
			s := &Signal{}
			ch <- s
		}
		close(ch)
		wg.Done()
	}()
	go func() {
		for v := range ch {
			_ = *v
		}
		wg.Done()
	}()
	wg.Wait()
}
