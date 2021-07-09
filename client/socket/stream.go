package socket

import (
	"alcor/auth"
	"alcor/client"
	"alcor/ws"
	"crypto/aes"
	"crypto/rand"
	"errors"

	"github.com/cloudflare/circl/dh/sidh"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

func EnrollClient(url string) error {
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

			block, err := aes.NewCipher(secret[:32])
			if err != nil {
				return err
			}

			pubkey := new(client.PublicKey)
			temp := make([]byte, len(account.PublicKey))
			for i := 0; i < len(account.PublicKey); i++ {
				temp[i] = account.PublicKey[i]
			}
			max := 0
			for max < len(temp) {
				max += block.BlockSize()
			}
			for max > len(temp) {
				temp = append(temp, 0)
			}
			encrypt := make([]byte, block.BlockSize())
			for len(temp) > 0 {
				block.Encrypt(encrypt, temp[:block.BlockSize()])
				temp = temp[block.BlockSize():]
			}
			pubkey.Key = encrypt
			bs, err := proto.Marshal(pubkey)
			if err != nil {
				return err
			}

			if err := conn.WriteMessage(websocket.BinaryMessage, bs); err != nil {
				return err
			}

			state = 2
		case 2:
			result := new(client.Result)
			if err := proto.Unmarshal(message, result); err != nil {
				return err
			}
			if !result.Ok {
				return errors.New(string(result.Error))
			}

			break loop
		}
	}

	return nil
}
