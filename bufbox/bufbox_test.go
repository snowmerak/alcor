package bufbox_test

import (
	"alcor/bufbox"
	"testing"
)

func TestMakePut(t *testing.T) {
	b := bufbox.NewWrapper().SetCompression(bufbox.BasisCompression).SetDirectory("./testDB").Unwrap()
	err := b.Put([]byte("aaa"), []byte("bbb"))
	if err != nil {
		t.Log(err)
	}
}
