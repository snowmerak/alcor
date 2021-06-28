package bufbox

type BufBox struct {
	directory         string
	compressionWriter func([]byte) ([]byte, error)
	compressionReader func([]byte) ([]byte, error)
}
