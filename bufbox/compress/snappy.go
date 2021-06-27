package compress

import (
	"bytes"

	"github.com/golang/snappy"
)

func WriteSnappy(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	w := snappy.NewBufferedWriter(buf)
	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Flush(); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func ReadSnappy(data []byte) ([]byte, error) {
	red := bytes.NewBuffer(data)
	r := snappy.NewReader(red)
	temp := make([]byte, 4096)
	buf := bytes.NewBuffer(nil)
	for {
		n, err := r.Read(temp)
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
