package admin

import (
	"alcor/db"
	"alcor/model/candidate"
	"alcor/model/capsule"
	"alcor/service/security"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/proto"
)

func MakeCheckPassword(password string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		capsule := new(capsule.Capsule)
		if err := proto.Unmarshal(c.Body(), capsule); err != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		data, err := security.Decrypt(capsule.ID, capsule.Data)
		if err != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		pw := string(data)
		if password == pw {
			return c.SendStatus(fasthttp.StatusOK)
		}
		return c.SendStatus(fasthttp.StatusBadRequest)
	}
}

func MakePostCandidate(password string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		capsule := new(capsule.Capsule)
		if err := proto.Unmarshal(c.Body(), capsule); err != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		data, err := security.Decrypt(capsule.ID, capsule.Data)
		if err != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		candidate := new(candidate.Candidate)
		if err := proto.Unmarshal(data, candidate); err != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		if candidate.Password != password {
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

func MakeDeleteCandidate(password string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		capsule := new(capsule.Capsule)
		if err := proto.Unmarshal(c.Body(), capsule); err != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		data, err := security.Decrypt(capsule.ID, capsule.Data)
		if err != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		candidate := new(candidate.Candidate)
		if err := proto.Unmarshal(data, candidate); err != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		if candidate.Password != password {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		candidatepb := new(db.Candidate)
		candidatepb.Name = candidate.Name

		if err := db.DeleteCandidate(context.Background(), candidatepb); err != nil {
			return c.SendStatus(fasthttp.StatusConflict)
		}

		return c.SendStatus(fasthttp.StatusOK)
	}
}

func MakePatchCandidate(password string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		capsule := new(capsule.Capsule)
		if err := proto.Unmarshal(c.Body(), capsule); err != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		data, err := security.Decrypt(capsule.ID, capsule.Data)
		if err != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		candidate := new(candidate.Candidate)
		if err := proto.Unmarshal(data, candidate); err != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		if candidate.Password != password {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		candidatepb := new(db.Candidate)
		candidatepb.Name = candidate.Name
		candidatepb.Markdown = candidate.Markdown

		if err := db.UpdateCandidate(context.Background(), candidatepb); err != nil {
			log.Println(err)
			return c.SendStatus(fasthttp.StatusConflict)
		}

		return c.SendStatus(fasthttp.StatusOK)
	}
}
