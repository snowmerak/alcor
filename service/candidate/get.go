package candidate

import (
	"alcor/db"
	"alcor/model/candidate"
	"context"
	"encoding/base64"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/proto"
)

func GetService(c *fiber.Ctx) error {
	candidates := new([]db.Candidate)
	if err := db.SelectCandidates(context.Background(), candidates); err != nil {
		c.SendStatus(fasthttp.StatusConflict)
		return c.Send([]byte(err.Error()))
	}

	list := new(candidate.List)
	for _, v := range *candidates {
		list.Names = append(list.Names, v.Name)
	}

	data, err := proto.Marshal(list)
	if err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusConflict)
		return c.Send([]byte(err.Error()))
	}

	c.SendStatus(fasthttp.StatusOK)
	return c.Send(data)
}

func GetAService(c *fiber.Ctx) error {
	cd := new(db.Candidate)
	body, err := base64.RawURLEncoding.DecodeString(c.Params("name", ""))
	if err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusBadRequest)
		return c.Send([]byte(err.Error()))
	}

	cd.Name = string(body)
	if err := db.SelectCandidate(context.Background(), cd); err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusConflict)
		return c.Send([]byte(err.Error()))
	}

	cand := new(candidate.Candidate)
	cand.Name = cd.Name
	cand.Markdown = cd.Markdown
	data, err := proto.Marshal(cand)
	if err != nil {
		log.Println(err)
		c.SendStatus(fasthttp.StatusConflict)
		return c.Send([]byte(err.Error()))
	}

	c.SendStatus(fasthttp.StatusOK)
	return c.Send(data)
}
