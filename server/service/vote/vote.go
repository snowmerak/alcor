package vote

import (
	"alcor/db"
	"alcor/service/auth"
	"alcor/service/vote/votepb"
	"alcor/service/wst"
	"alcor/worker/hashbuffer"
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"math/big"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/snowmerak/jumper"
	"github.com/snowmerak/wrapper/buffer"
	"google.golang.org/protobuf/proto"
)

type VoteError struct {
	errorData string
}

func (v VoteError) Error() string {
	return v.errorData
}

func NewVoteError(data string) VoteError {
	return VoteError{errorData: data}
}

func PhoneCertSocket(c *fiber.Ctx) error {
	return wst.Listen(c.Context(), func(conn *websocket.Conn) {
		buf := bytes.NewBuffer(nil)
		aead, err := wst.Handshake(conn, buf)
		if err != nil {
			jumper.Offer(NewVoteError(err.Error()))
			return
		}

		secret := make([]byte, buf.Len())
		if err := buffer.BytesBuffer2Bytes(secret, buf); err != nil {
			jumper.Offer(NewVoteError(err.Error()))
			return
		}

		const randomBytesLength = 64
		randomBytes := make([]byte, randomBytesLength)
		if n, err := rand.Read(randomBytes); n < randomBytesLength || err != nil {
			if err != nil {
				jumper.Offer(NewVoteError(err.Error()))
				return
			}
			jumper.Offer(NewVoteError("cannot read all bytes"))
			return
		}
		if err := wst.SendBinaryMessage(conn, aead, secret, randomBytes); err != nil {
			jumper.Offer(NewVoteError(err.Error()))
			return
		}

		errmsg := ""

		if errmsg == "" {
			buf.Reset()
			if err := wst.ReceiveBinaryMessage(conn, aead, secret, buf); err != nil {
				jumper.Offer(NewVoteError(err.Error()))
				errmsg = err.Error()
			}
		}

		paper := new(votepb.Paper)
		if errmsg == "" {
			if err := proto.Unmarshal(buf.Bytes(), paper); err != nil {
				jumper.Offer(NewVoteError(err.Error()))
				errmsg = err.Error()
			}
		}

		voter := new(db.Voter)
		if errmsg == "" {
			voter.HashID = paper.VoterId
			if err := db.SelectVoter(context.Background(), voter); err != nil {
				jumper.Offer(NewVoteError(err.Error()))
				errmsg = err.Error()
			}
		}

		if errmsg == "" {
			if voter.Voted {
				jumper.Offer(NewVoteError("already voted"))
				errmsg = "already voted"
			}
		}

		r, s := new(big.Int).SetInt64(int64(0)), new(big.Int).SetInt64(int64(0))
		if errmsg == "" {
			r, s, err = auth.DeserializeSignature(paper.Signature)
			if err != nil {
				jumper.Offer(NewVoteError(err.Error()))
				errmsg = err.Error()
			}
		}

		publickey, err := auth.DeserializePublicKey(voter.PublicKey)
		if errmsg == "" {
			if err != nil {
				jumper.Offer(NewVoteError(err.Error()))
				errmsg = err.Error()
			}
		}

		if errmsg == "" {
			sha := sha512.New()
			sha.Write(paper.RandomBytes)
			sha.Write(paper.VoterId)
			sha.Write(paper.Timestamp)
			sha.Write(paper.Message)
			hash := sha.Sum(nil)
			if !bytes.Equal(paper.Hash, hash) {
				jumper.Offer(NewVoteError("paper hash is invalid"))
				errmsg = "paper hash is invalid"
			}
			if errmsg == "" && !ecdsa.Verify(publickey, paper.Hash, r, s) {
				jumper.Offer(NewVoteError("paper hash is invalid"))
				errmsg = "paper hash is invalid"
			}
		}

		if errmsg == "" {
			paperTx := new(db.Paper)
			paperTx.VoterID = paper.VoterId
			paperTx.RandomBytes = paper.RandomBytes
			paperTx.Message = paper.Message
			paperTx.Timestamp = paper.Timestamp
			paperTx.Signature = paper.Signature
			paperTx.Hash = paper.Hash
			if err := db.InsertPaper(context.Background(), paperTx); err != nil {
				jumper.Offer(NewVoteError(err.Error()))
				errmsg = err.Error()
			}
		}

		if errmsg == "" {
			voter.RandomBytes = paper.RandomBytes
			voter.Voted = true
			voter.Timestamp = paper.Timestamp
			if err := db.UpdateVoter(context.Background(), voter); err != nil {
				jumper.Offer(NewVoteError(err.Error()))
				errmsg = err.Error()
			}
		}

		if errmsg == "" {
			hashbuffer.EnQueue(paper.Hash)
		}

		result := new(votepb.Result)
		result.Error = errmsg
		result.Ok = true
		if errmsg != "" {
			result.Ok = false
		}
		data, err := proto.Marshal(result)
		if err != nil {
			jumper.Offer(NewVoteError(err.Error()))
			return
		}
		if err := wst.SendBinaryMessage(conn, aead, secret, data); err != nil {
			jumper.Offer(NewVoteError(err.Error()))
			return
		}

		jumper.Offer(NewVoteError(fmt.Sprintf("%x's %x is added", paper.VoterId, paper.Hash)))
	})
}
