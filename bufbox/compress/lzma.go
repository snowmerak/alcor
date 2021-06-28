package compress

import (
	"bytes"

	"github.com/ulikunitz/xz/lzma"
)

func WriteLZMA2(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	w, err := lzma.NewWriter2(buf)
	if err != nil {
		return nil, err
	}
	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ReadLZMA2(data []byte) ([]byte, error) {
	red := bytes.NewBuffer(data)
	r, err := lzma.NewReader2(red)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	temp := make([]byte, 4096)
	for !r.EOS() {
		n, err := r.Read(temp)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
		buf.Write(temp[:n])
	}
	return buf.Bytes(), nil
}

func LZMA2() (func([]byte) ([]byte, error), func([]byte) ([]byte, error)) {
	return WriteLZMA2, ReadLZMA2
}
