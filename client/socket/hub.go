package socket

import (
	"alcor/client"
	"alcor/client/book"
	"alcor/ws"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
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

	{
		result := new(client.Result)
		result.Ok = true
		result.Error = nil
		bs, err := proto.Marshal(result)
		if err != nil {
			return err
		}
		conn.WriteMessage(websocket.BinaryMessage, bs)
	}

	hashed := sha256.Sum256(secert)
	block, err := aes.NewCipher(hashed[:])
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

		aead, err := cipher.NewGCM(block)
		if err != nil {
			return err
		}

		c.PublicKey, err = aead.Open(nil, secert[:aead.NonceSize()], c.PublicKey, nil)
		if err != nil {
			return err
		}
	}

	rs := new(client.HashID)
	salt := make([]byte, 128)
	rand.Read(salt)
	salt = append(salt, c.ID...)
	id := sha512.Sum512(salt)
	rs.ID = id[:]
	c.ID = id[:]
	rs.Error = nil
	if err := book.Put(c); err != nil {
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
