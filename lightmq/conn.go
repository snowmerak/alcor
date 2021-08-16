package lightmq

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/fasthttp/websocket"
)

const Uri = "lightmq"

func Subscript(uri, mainCategory, subCategory string) (*websocket.Conn, error) {
	conn, _, err := websocket.DefaultDialer.Dial("ws://"+uri+":11211/"+mainCategory+"/"+subCategory, nil)
	return conn, err
}

func ReadMessage(conn *websocket.Conn) ([]byte, error) {
	var typ = websocket.TextMessage
	var data = []byte{}
	var err = error(nil)
	for typ != websocket.BinaryMessage {
		typ, data, err = conn.ReadMessage()
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func Publish(uri, mainCategory, subCategory string, data []byte) error {
	res, err := http.Post("http://"+uri+":11211/"+mainCategory+"/"+subCategory, "application/octet-stream", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	if res.StatusCode == http.StatusBadRequest {
		return errors.New("request is bad")
	}
	return nil
}
