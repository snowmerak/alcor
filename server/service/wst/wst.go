package wst

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"github.com/aead/ecdh"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

func Handshake(c *websocket.Conn, writer io.Writer) (cipher.AEAD, error) {
	_, data, err := c.ReadMessage()
	if err != nil {
		return nil, err
	}

	ec := ecdh.X25519()
	priv, pub, err := ec.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	p := pub.([32]uint8)
	if err := c.WriteMessage(websocket.BinaryMessage, p[:]); err != nil {
		return nil, err
	}

	sec := ec.ComputeSecret(priv, data)
	block, err := aes.NewCipher(sec)
	if err != nil {
		return nil, err
	}
	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if _, err := writer.Write(sec); err != nil {
		return nil, err
	}

	return aead, nil
}

var Upgrader = websocket.FastHTTPUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(ctx *fasthttp.RequestCtx) bool { return true },
}

func Listen(ctx *fasthttp.RequestCtx, handler func(conn *websocket.Conn)) error {
	return Upgrader.Upgrade(ctx, handler)
}

func SendBinaryMessage(conn *websocket.Conn, aead cipher.AEAD, secret, data []byte) error {
	encrypted := aead.Seal(nil, secret[:aead.NonceSize()], data, nil)
	return conn.WriteMessage(websocket.BinaryMessage, encrypted)
}

func ReceiveBinaryMessage(conn *websocket.Conn, aead cipher.AEAD, secret []byte, writer io.Writer) error {
	typ, data, err := conn.ReadMessage()
	if err != nil {
		return err
	}
	if typ != websocket.BinaryMessage {
		return NotBinaryMessage()
	}
	plain, err := aead.Open(nil, secret[:aead.NonceSize()], data, nil)
	if err != nil {
		return err
	}
	_, err = writer.Write(plain)
	return err
}

func NotBinaryMessage() error {
	return errors.New("message is not binary message")
}
