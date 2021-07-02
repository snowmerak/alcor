package bufbox

import (
	"os"
	"path/filepath"
	"strings"
)

func Mkdir(dir string) error {
	dirs := strings.Split(dir, string(filepath.Separator))
	for i := 0; i < len(dirs); i++ {
		path := filepath.Join(dirs[:i+1]...)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			if err := os.Mkdir(path, 0700); err != nil {
				return err
			}
		}
	}
	return nil
}
