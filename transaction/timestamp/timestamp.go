package timestamp

import (
	"encoding/binary"
	"time"
)

func Now() []byte {
	t := time.Now().UnixNano()
	bs := make([]byte, 64)
	binary.BigEndian.PutUint64(bs, uint64(t))
	return bs
}
