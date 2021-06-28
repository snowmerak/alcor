package compress

import (
	"bytes"

	"github.com/andybalholm/brotli"
)

func WriteBrotli(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	brt := brotli.NewWriterLevel(buf, brotli.BestCompression)
	if _, err := brt.Write(data); err != nil {
		return nil, err
	}
	if err := brt.Flush(); err != nil {
		return nil, err
	}
	if err := brt.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ReadBrotli(data []byte) ([]byte, error) {
	red := bytes.NewBuffer(data)
	brt := brotli.NewReader(red)
	buf := bytes.NewBuffer(nil)
	temp := make([]byte, 4096)
	for {
		n, err := brt.Read(temp)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
		if n == 0 {
			break
		}
		buf.Write(temp[:n])
	}
	return buf.Bytes(), nil
}

func Brotli() (func([]byte) ([]byte, error), func([]byte) ([]byte, error)) {
	return WriteBrotli, ReadBrotli
}
