package admin

import (
	"alcor/db"
	"alcor/model/candidate"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/proto"
)

func MakePostCandidate(password string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		candidate := new(candidate.Candidate)
		if err := proto.Unmarshal(c.Body(), candidate); err != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		if candidate.Password == password {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		candidatepb := new(db.Candidate)
		candidatepb.Name = candidate.Name
		candidatepb.Markdown = candidate.Markdown

		if err := db.InsertCandidate(context.Background(), candidatepb); err != nil {
			return c.SendStatus(fasthttp.StatusConflict)
		}

		return c.SendStatus(fasthttp.StatusOK)
	}
}
