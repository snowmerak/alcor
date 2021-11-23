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
	id, err := base64.URLEncoding.DecodeString(c.Params("id", ""))
	if err != nil {
		return c.SendStatus(fasthttp.StatusBadRequest)
	}
	voter := new(db.Voter)
	voter.HashID = id
	if err := db.SelectVoter(context.Background(), voter); err != nil {
		return c.SendStatus(fasthttp.StatusNotFound)
	}

	if !voter.Voted {
		return c.SendStatus(fasthttp.StatusBadRequest)
	}

	reply := new(info.Voter)
	reply.HashID = voter.HashID
	reply.RandomBytes = voter.RandomBytes
	reply.PublicKey = voter.PublicKey
	reply.Timestamp = voter.Timestamp
	reply.Voted = voter.Voted
	reply.IsCandidate = false
	reply.RepliedSurvey = false
	data, err := proto.Marshal(reply)
	if err != nil {
		return c.SendStatus(fasthttp.StatusConflict)
	}

	c.SendStatus(fasthttp.StatusOK)
	return c.Send(data)
}
