package security

import (
	"alcor/model/security"
	"alcor/worker/sessions"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	"encoding/hex"
	"log"

	"github.com/aead/ecdh"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/proto"
)

func Handshake(data []byte) ([]byte, []byte, error) {
	ec := ecdh.X25519()
	priv, pub, err := ec.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	p := pub.([32]uint8)

	sec := ec.ComputeSecret(priv, data)

	return sec, p[:], nil
}

func Secret2Aead(secret []byte) (cipher.AEAD, error) {
	block, err := aes.NewCipher(secret[:32])
	if err != nil {
		return nil, err
	}
	return cipher.NewGCM(block)
}

func Service(c *fiber.Ctx) error {
	secret, pub, err := Handshake(c.Body())
	if err != nil {
		log.Println(err)
		c.Response().Header.SetStatusCode(fasthttp.StatusBadRequest)
		return c.Send([]byte(err.Error()))
	}
	hashed := sha512.Sum512(secret)
	reply := new(security.Reply)
	reply.ID = hex.EncodeToString(hashed[:])
	reply.PublicKey = pub
	sessions.Set(reply.ID, secret)
	data, err := proto.Marshal(reply)
	if err != nil {
		log.Println(err)
		c.Response().Header.SetStatusCode(fasthttp.StatusNotModified)
		return c.Send([]byte(err.Error()))
	}
	return c.Send(data)
}
