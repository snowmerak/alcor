package socket

import (
	"alcor/client"
	"alcor/client/book"
	"alcor/ws"
	"crypto/aes"
	"crypto/rand"
	"net/http"

	"github.com/cloudflare/circl/dh/sidh"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func EnrollServer(rw http.ResponseWriter, r *http.Request) error {
	conn, err := ws.Upgrader.Upgrade(rw, r, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	privateKey := sidh.NewPrivateKey(sidh.Fp503, sidh.KeyVariantSidhA)
	if err := privateKey.Generate(rand.Reader); err != nil {
		return err
	}
	publicKey := sidh.NewPublicKey(sidh.Fp503, sidh.KeyVariantSidhA)
	privateKey.GeneratePublicKey(publicKey)

	{
		pubKey := new(client.PublicKey)
		pubKey.Key = make([]byte, publicKey.Size())
		publicKey.Export(pubKey.Key)
		bs, err := proto.Marshal(pubKey)
		if err != nil {
			return err
		}
		if err := conn.WriteMessage(websocket.BinaryMessage, bs); err != nil {
			return err
		}
	}

	{
		typ, message, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		if typ != websocket.BinaryMessage {
			return ws.NotBinaryMessage()
		}
		pubKey := new(client.PublicKey)
		if err := proto.Unmarshal(message, pubKey); err != nil {
			return err
		}
		publicKey = sidh.NewPublicKey(sidh.Fp503, sidh.KeyVariantSidhB)
		if err := publicKey.Import(pubKey.Key); err != nil {
			return err
		}
	}

	secert := make([]byte, privateKey.SharedSecretSize())
	privateKey.DeriveSecret(secert, publicKey)

	block, err := aes.NewCipher(secert[:32])
	if err != nil {
		return err
	}

	c := new(client.Client)
	{
		typ, message, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		if typ != websocket.BinaryMessage {
			return ws.NotBinaryMessage()
		}
		if err := proto.Unmarshal(message, c); err != nil {
			return err
		}
		temp := c.PublicKey
		c.PublicKey = make([]byte, len(temp))
		max := 0
		for max < len(temp) {
			max += block.BlockSize()
		}
		for max > len(temp) {
			temp = append(temp, 0)
		}
		decrypted := make([]byte, block.BlockSize())
		for len(temp) > 0 {
			block.Decrypt(decrypted, temp[:block.BlockSize()])
			c.PublicKey = append(c.PublicKey, decrypted...)
			temp = temp[block.BlockSize():]
		}
	}

	rs := new(client.Result)
	rs.Ok = true
	rs.Error = nil
	if err := book.Put(c); err != nil {
		rs.Ok = false
		rs.Error = []byte(err.Error())
	}

	bs, err := proto.Marshal(rs)
	if err != nil {
		return err
	}

	if err := conn.WriteMessage(websocket.BinaryMessage, bs); err != nil {
		return err
	}

	return nil
}
