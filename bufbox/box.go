package bufbox

import (
	"alcor/bufbox/compress"
	"alcor/bufbox/csset"
	"encoding/hex"
	"errors"
	"os"
	"path/filepath"
)

type BufBox struct {
	directory         string
	compressionWriter func([]byte) ([]byte, error)
	compressionReader func([]byte) ([]byte, error)
	cacheSet          csset.CSSet
}

var BasisCompression = compress.Brotli

func (b *BufBox) nameToPath(name []byte) string {
	return filepath.Join(b.directory, hex.EncodeToString(name))
}

func (b *BufBox) Exist(name []byte) bool {
	path := b.nameToPath(name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func (b *BufBox) Get(name []byte) ([]byte, error) {
	content, err := os.ReadFile(b.nameToPath(name))
	if err != nil {
		return nil, err
	}
	if b.compressionReader != nil {
		content, err = b.compressionReader(content)
		if err != nil {
			return nil, err
		}
	}
	return content, nil
}

func (b *BufBox) Put(name []byte, content []byte) error {
	if b.Exist(name) {
		return errors.New("BufBox: already key")
	}
	if b.compressionWriter != nil {
		err := error(nil)
		content, err = b.compressionWriter(content)
		if err != nil {
			return err
		}
	}
	return os.WriteFile(b.nameToPath(name), content, 0o770)
}

func (b *BufBox) List() ([][]byte, error) {
	list, err := os.ReadDir(b.directory)
	if err != nil {
		return nil, err
	}
	rs := [][]byte{}
	for _, l := range list {
		bs, err := hex.DecodeString(l.Name())
		if err != nil {
			continue
		}
		rs = append(rs, bs)
	}
	return rs, nil
}
