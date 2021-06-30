package bufbox

import "alcor/bufbox/csset"

type BufBoxWrapper struct {
	bufBox *BufBox
}

func (b *BufBoxWrapper) Unwrap() *BufBox {
	b.bufBox.cacheSet = csset.CSSet{}
	return b.bufBox
}

func (b *BufBoxWrapper) SetCompression(comp func() (func([]byte) ([]byte, error), func([]byte) ([]byte, error))) *BufBoxWrapper {
	b.bufBox.compressionWriter, b.bufBox.compressionReader = comp()
	return b
}

func (b *BufBoxWrapper) SetDirectory(dir string) *BufBoxWrapper {
	b.bufBox.directory = dir
	return b
}
