package info

import (
	"alcor/db"
	"alcor/model/info"
	"context"
	"encoding/base64"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/proto"
)

func GetService(c *fiber.Ctx) error {
	id, err := base64.RawURLEncoding.DecodeString(c.Params("id", ""))
	if err != nil {
		return c.SendStatus(fasthttp.StatusBadRequest)
	}
	paper := new(db.Paper)
	paper.Hash = id
	if err := db.SelectPaper(context.Background(), paper); err != nil {
		return c.SendStatus(fasthttp.StatusNotFound)
	}

	reply := new(info.Paper)
	reply.Hash = paper.Hash
	reply.Message = paper.Message
	reply.RandomBytes = paper.RandomBytes
	reply.Signature = paper.Signature
	reply.Timestamp = paper.Timestamp
	reply.VoterID = paper.VoterID
	data, err := proto.Marshal(reply)
	if err != nil {
		return c.SendStatus(fasthttp.StatusConflict)
	}

	c.SendStatus(fasthttp.StatusOK)
	return c.Send(data)
}
