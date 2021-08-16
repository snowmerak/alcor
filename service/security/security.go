package security

import (
	"alcor/model/capsule"
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

func Encrypt(id string, data []byte) ([]byte, error) {
	secret, err := sessions.Get(id)
	if err != nil {
		return nil, err
	}
	aead, err := Secret2Aead(secret)
	if err != nil {
		return nil, err
	}
	return aead.Seal(nil, secret[:aead.NonceSize()], data, nil), nil
}

func Decrypt(id string, data []byte) ([]byte, error) {
	secret, err := sessions.Get(id)
	if err != nil {
		return nil, err
	}
	aead, err := Secret2Aead(secret)
	if err != nil {
		return nil, err
	}
	return aead.Open(nil, secret[:aead.NonceSize()], data, nil)
}

func Encapsulate(id string, data []byte) ([]byte, error) {
	content, err := Encrypt(id, data)
	if err != nil {
		return nil, err
	}
	capsule := new(capsule.Capsule)
	capsule.ID = id
	capsule.Data = content
	return proto.Marshal(capsule)
}

func Decapsulate(data []byte) (string, []byte, error) {
	capsule := new(capsule.Capsule)
	if err := proto.Unmarshal(data, capsule); err != nil {
		return "", nil, err
	}
	data, err := Decrypt(capsule.ID, capsule.Data)
	return capsule.ID, data, err
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
