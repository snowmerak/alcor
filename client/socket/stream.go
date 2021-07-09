package socket

import (
	"alcor/auth"
	"alcor/client"
	"alcor/client/wallet"
	"alcor/ws"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"

	"github.com/cloudflare/circl/dh/sidh"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func EnrollClient(url string, id string) error {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	sidhPrivKey := sidh.NewPrivateKey(sidh.Fp503, sidh.KeyVariantSidhB)
	sidhPrivKey.Generate(rand.Reader)
	sidhPubKey := sidh.NewPublicKey(sidh.Fp503, sidh.KeyVariantSidhB)
	sidhPrivKey.GeneratePublicKey(sidhPubKey)
	secret := make([]byte, sidhPrivKey.SharedSecretSize())

	privateKey, err := auth.Basis.GeneratePrivateKey()
	if err != nil {
		return err
	}
	publicKey, err := auth.Basis.GeneratePublicKey(privateKey)
	if err != nil {
		return err
	}

	account := new(client.Account)
	account.PrivateKey, err = auth.Basis.SerializePrivateKey(privateKey)
	if err != nil {
		return err
	}
	account.PublicKey, err = auth.Basis.SerializePublicKey(publicKey)
	if err != nil {
		return err
	}

	state := 0
loop:
	for {
		typ, message, err := conn.ReadMessage()
		if err != nil {
			return err
		}
		if typ != websocket.BinaryMessage {
			return ws.NotBinaryMessage()
		}
		switch state {
		case 0:
			pubkey := new(client.PublicKey)
			if err := proto.Unmarshal(message, pubkey); err != nil {
				return err
			}

			data := make([]byte, sidhPubKey.Size())
			sidhPubKey.Export(data)

			sidhPubKey = sidh.NewPublicKey(sidh.Fp503, sidh.KeyVariantSidhA)
			if err := sidhPubKey.Import(pubkey.Key); err != nil {
				return err
			}

			pubkey.Key = data
			data, err = proto.Marshal(pubkey)
			if err != nil {
				return err
			}

			if err := conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
				return err
			}

			sidhPrivKey.DeriveSecret(secret, sidhPubKey)

			state = 1
		case 1:
			result := new(client.Result)
			if err := proto.Unmarshal(message, result); err != nil {
				return err
			}
			if !result.Ok {
				return errors.New(string(result.Error))
			}

			hashed := sha256.Sum256(secret)
			block, err := aes.NewCipher(hashed[:])
			if err != nil {
				return err
			}

			pubkey := new(client.PublicKey)

			aead, err := cipher.NewGCM(block)
			if err != nil {
				return err
			}
			pubkey.Key = aead.Seal(nil, secret[:aead.NonceSize()], account.PublicKey, nil)

			c := new(client.Client)
			c.ID = []byte(id)
			c.PublicKey = pubkey.Key
			bs, err := proto.Marshal(c)
			if err != nil {
				return err
			}

			if err := conn.WriteMessage(websocket.BinaryMessage, bs); err != nil {
				return err
			}

			state = 2
		case 2:
			result := new(client.HashID)
			if err := proto.Unmarshal(message, result); err != nil {
				return err
			}
			if result.Error != nil {
				return errors.New(string(result.Error))
			} else {
				account.ID = result.ID
			}

			break loop
		}
	}

	if err := wallet.Put(account); err != nil {
		return err
	}

	return nil
}
