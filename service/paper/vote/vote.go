package vote

import (
	"alcor/db"
	"alcor/model/vote"
	"alcor/service/auth"
	"alcor/service/security"
	"alcor/worker/hashbuffer"
	"alcor/worker/keybytes"
	"bytes"
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/proto"
)

func RequestService(c *fiber.Ctx) error {
	id, data, err := security.Decapsulate(c.Body())
	if err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusBadRequest)
		return c.Send([]byte(err.Error()))
	}

	voter := new(db.Voter)
	voter.HashID = data
	if err := db.SelectVoter(context.Background(), voter); err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusConflict)
		return c.Send([]byte(err.Error()))
	}

	if voter.Voted {
		log.Println(err)
		c.SendStatus(fasthttp.StatusBadRequest)
		return c.Send([]byte(err.Error()))
	}

	key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusConflict)
		return c.Send([]byte(err.Error()))
	}

	if err := keybytes.Set(id, key); err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusConflict)
		return c.Send([]byte(err.Error()))
	}

	data, err = security.Encapsulate(id, key)
	if err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusConflict)
		return c.Send([]byte(err.Error()))
	}

	c.SendStatus(fasthttp.StatusOK)
	return c.Send(data)
}

func SubmitService(c *fiber.Ctx) error {
	id, data, err := security.Decapsulate(c.Body())
	if err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusBadRequest)
		return c.Send([]byte(err.Error()))
	}

	p := new(vote.Paper)
	if err := proto.Unmarshal(data, p); err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusBadRequest)
		return c.Send([]byte(err.Error()))
	}

	key, err := keybytes.Get(id)
	if err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusNotFound)
		return c.Send([]byte(err.Error()))
	}

	if !bytes.Equal(key, p.RandomBytes) {
		log.Println("keybytes is invalid")
		c.SendStatus(fasthttp.StatusBadRequest)
		return c.Send([]byte("keybytes is invalid"))
	}

	voter := new(db.Voter)
	voter.HashID = p.VoterId
	if err := db.SelectVoter(context.Background(), voter); err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusNotFound)
		return c.Send([]byte(err.Error()))
	}

	if voter.Voted {
		log.Println(err)
		c.SendStatus(fasthttp.StatusBadRequest)
		return c.Send([]byte("already voted"))
	}

	sha := sha256.New()
	sha.Write(p.RandomBytes)
	sha.Write(p.VoterId)
	sha.Write(p.Timestamp)
	sha.Write(p.Message)
	hashed := sha.Sum(nil)
	if !bytes.Equal(p.Hash, hashed) {
		log.Println(err)
		c.SendStatus(fasthttp.StatusBadRequest)
		return c.Send([]byte(err.Error()))
	}

	publicKey, err := auth.DeserializePublicKey(voter.PublicKey)
	if err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusBadRequest)
		return c.Send([]byte(err.Error()))
	}

	r, s, err := auth.DeserializeSignature(p.Signature)
	if err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusBadRequest)
		return c.Send([]byte(err.Error()))
	}

	if !ecdsa.Verify(publicKey, p.Hash, r, s) {
		log.Println(err)
		c.SendStatus(fasthttp.StatusBadRequest)
		return c.Send([]byte("signatue is invalid"))
	}

	paper := new(db.Paper)
	paper.VoterID = p.VoterId
	paper.Message = p.Message
	paper.RandomBytes = p.RandomBytes
	paper.Timestamp = p.Timestamp
	paper.Signature = p.Signature
	paper.Hash = p.Hash
	if err := db.InsertPaper(context.Background(), paper); err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusConflict)
		return c.Send([]byte(err.Error()))
	}

	go func() { hashbuffer.EnQueue(paper.Hash) }()

	c.SendStatus(fasthttp.StatusOK)

	voter.Voted = true
	voter.Timestamp = paper.Timestamp
	voter.RandomBytes = paper.RandomBytes
	if err := db.UpdateVoter(context.Background(), voter); err != nil {
		log.Println(err)
	}

	data, err = security.Encapsulate(id, paper.Hash)
	if err != nil {
		log.Println(err)
	}

	return c.Send(data)
}
