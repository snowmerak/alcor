package socket

import (
	"alcor/auth"
	"alcor/client"
	"alcor/transaction"
	"errors"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func Connect(url string, account *client.Account, message string) (*transaction.Result, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	state := 0

	for {
		typ, msg, err := conn.ReadMessage()
		if err != nil {
			return nil, err
		}
		if typ != websocket.BinaryMessage {
			return nil, errors.New("data is not binary message")
		}
		/*
			0: receive Key date and make Transaction then
			send Transaction and disconnect
			1: receive ok or error
		*/
		switch state {
		case 0:
			key := new(transaction.Key)
			if err := proto.Unmarshal(msg, key); err != nil {
				return nil, err
			}
			t := new(transaction.Transaction)
			t.From = account.ID
			t.Key = key.Key
			t.Message = []byte(message)
			t.Timestamp = key.Timestamp
			t.Hash = transaction.MakeHash(t)
			privateKey, err := auth.Basis.DeserializePrivateKey(account.PrivateKey)
			if err != nil {
				return nil, err
			}
			r, s, err := auth.Basis.Sign(privateKey, t.Hash)
			if err != nil {
				return nil, err
			}
			sign, err := auth.Basis.SerializeSignature(r, s)
			if err != nil {
				return nil, err
			}
			t.Signature = sign
			data, err := proto.Marshal(t)
			if err != nil {
				return nil, err
			}
			conn.WriteMessage(websocket.BinaryMessage, data)
			state = 1
		case 1:
			result := new(transaction.Result)
			if err := proto.Unmarshal(msg, result); err != nil {
				return nil, err
			}
			return result, conn.Close()
		}
	}
}
