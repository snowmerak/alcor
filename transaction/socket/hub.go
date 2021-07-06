package socket

import (
	"alcor/auth"
	"alcor/client/book"
	"alcor/transaction"
	"alcor/transaction/ground"
	"alcor/transaction/timestamp"
	"bytes"
	"errors"
	"math/rand"
	"net/http"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func Handle(rw http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(rw, r, nil)
	if err != nil {
		return err
	}

	key := new(transaction.Key)
	key.Key = make([]byte, 64)
	rand.Read(key.Key)
	key.Timestamp = timestamp.Now()
	bs, err := proto.Marshal(key)
	if err != nil {
		return err
	}
	if err := conn.WriteMessage(websocket.BinaryMessage, bs); err != nil {
		return err
	}

	typ, msg, err := conn.ReadMessage()
	if err != nil {
		return err
	}
	if typ != websocket.BinaryMessage {
		return errors.New("data is not binary message")
	}

	t := new(transaction.Transaction)
	if err := proto.Unmarshal(msg, t); err != nil {
		return err
	}

	if !bytes.Equal(transaction.MakeHash(t), t.Hash) {
		return errors.New("hash is invalid")
	}

	if c, err := book.Get(t.From); err != nil {
		return err
	} else {
		publicKey, err := auth.Basis.DeserializePublicKey(c.PublicKey)
		if err != nil {
			return err
		}
		r, s, err := auth.Basis.DeserializeSignature(t.Signature)
		if err != nil {
			return err
		}
		if !auth.Basis.Verify(publicKey, t.Hash, r, s) {
			return errors.New("signature is invalid")
		}
	}

	return ground.Put(t)
}
