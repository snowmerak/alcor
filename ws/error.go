package ws

import "errors"

func NotBinaryMessage() error {
	return errors.New("message is not binary message")
}
