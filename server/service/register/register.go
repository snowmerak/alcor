package register

import (
	"alcor/db"
	"alcor/service/auth"
	"alcor/service/register/registerpb"
	"alcor/service/wst"
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/snowmerak/jumper"
	"github.com/snowmerak/wrapper/buffer"
	"google.golang.org/protobuf/proto"
)

type RegisterError struct {
	errorData string
}

func (v RegisterError) Error() string {
	return v.errorData
}

func NewRegisterError(data string) RegisterError {
	return RegisterError{errorData: data}
}

func PhoneRegisterSocket(c *fiber.Ctx) error {
	return wst.Listen(c.Context(), func(conn *websocket.Conn) {
		buf := bytes.NewBuffer(nil)

		aead, err := wst.Handshake(conn, buf)
		if err != nil {
			jumper.Offer(NewRegisterError(err.Error()))
			return
		}
		secret := make([]byte, buf.Len())
		if err := buffer.BytesBuffer2Bytes(secret, buf); err != nil {
			jumper.Offer(NewRegisterError(err.Error()))
			return
		}

		keyBytes := make([]byte, 64)
		if n, err := rand.Read(keyBytes); n < 64 || err != nil {
			if err != nil {
				jumper.Offer(NewRegisterError(err.Error()))
				return
			}
			jumper.Offer(NewRegisterError("cannot read all random bytes"))
			return
		}
		if err := wst.SendBinaryMessage(conn, aead, secret, keyBytes); err != nil {
			jumper.Offer(NewRegisterError(err.Error()))
			return
		}

		buf.Reset()
		if err := wst.ReceiveBinaryMessage(conn, aead, secret, buf); err != nil {
			jumper.Offer(NewRegisterError(err.Error()))
			return
		}

		errmsg := ""

		request := new(registerpb.Request)
		if errmsg == "" {
			if err := proto.Unmarshal(buf.Bytes(), request); err != nil {
				jumper.Offer(NewRegisterError(err.Error()))
				errmsg = err.Error()
			}
		}

		if errmsg == "" {
			if !bytes.Equal(keyBytes, request.KeyBytes) {
				jumper.Offer(NewRegisterError("key bytes is not equal"))
				errmsg = "key bytes is not equal"
			}
		}

		voter := new(db.Voter)
		voter.HashID = request.VoterID
		if errmsg == "" {
			if err := db.SelectVoter(context.Background(), voter); err != nil {
				jumper.Offer(NewRegisterError(err.Error()))
				errmsg = err.Error()
			}
		}

		if errmsg == "" {
			if voter.IsCandidate {
				jumper.Offer(NewRegisterError(err.Error()))
				errmsg = "aleady registered in candidates"
			}
		}

		publicKey := new(ecdsa.PublicKey)
		if errmsg == "" {
			publicKey, err = auth.DeserializePublicKey(voter.PublicKey)
			if err != nil {
				jumper.Offer(NewRegisterError(err.Error()))
				errmsg = err.Error()
			}
		}

		r, s := new(big.Int).SetInt64(0), new(big.Int).SetInt64(0)
		if errmsg == "" {
			r, s, err = auth.DeserializeSignature(request.Signature)
			if err != nil {
				jumper.Offer(NewRegisterError(err.Error()))
				errmsg = err.Error()
			}
		}

		if errmsg == "" {
			if !ecdsa.Verify(publicKey, request.Hash, r, s) {
				jumper.Offer(NewRegisterError("request's signature is invalid"))
				errmsg = "request's signature is invalid"
			}
		}

		candidate := new(db.Candidate)
		if errmsg == "" {
			candidate.Name = request.Name
			candidate.Markdown = request.Markdown
			if err := db.InsertCandidate(context.Background(), candidate); err != nil {
				jumper.Offer(NewRegisterError(err.Error()))
				errmsg = err.Error()
			}
		}

		if errmsg == "" {
			if err := db.ChangeVoter2Candidate(context.Background(), voter); err != nil {
				jumper.Offer(NewRegisterError(err.Error()))
				errmsg = err.Error()
			}
		}

		result := new(registerpb.Result)
		result.OK = true
		result.Error = errmsg
		if errmsg != "" {
			result.OK = false
		}
		data, err := proto.Marshal(result)
		if err != nil {
			jumper.Offer(NewRegisterError(err.Error()))
			return
		}
		if err := wst.SendBinaryMessage(conn, aead, secret, data); err != nil {
			jumper.Offer(NewRegisterError(err.Error()))
			return
		}

		if result.OK {
			jumper.Offer(NewRegisterError(fmt.Sprintf("%v is enrolled", request.Name)))
		}
	})
}
