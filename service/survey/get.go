package survey

import (
	"alcor/model/stats"
	"alcor/static"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/proto"
)

func GetRegions(c *fiber.Ctx) error {
	list := new(stats.List)
	list.Data = static.Regions

	data, err := proto.Marshal(list)
	if err != nil {
		return c.SendStatus(fasthttp.StatusConflict)
	}

	c.SendStatus(fasthttp.StatusOK)
	return c.Send(data)
}

func GetJobs(c *fiber.Ctx) error {
	list := new(stats.List)
	list.Data = static.Jobs

	data, err := proto.Marshal(list)
	if err != nil {
		return c.SendStatus(fasthttp.StatusConflict)
	}

	c.SendStatus(fasthttp.StatusOK)
	return c.Send(data)
}

func GetEducations(c *fiber.Ctx) error {
	list := new(stats.List)
	list.Data = static.Educations

	data, err := proto.Marshal(list)
	if err != nil {
		return c.SendStatus(fasthttp.StatusConflict)
	}

	c.SendStatus(fasthttp.StatusOK)
	return c.Send(data)
}

func GetsHouseTypes(c *fiber.Ctx) error {
	list := new(stats.List)
	list.Data = static.HouseTypes

	data, err := proto.Marshal(list)
	if err != nil {
		return c.SendStatus(fasthttp.StatusConflict)
	}

	c.SendStatus(fasthttp.StatusOK)
	return c.Send(data)
}

func GetIdeologys(c *fiber.Ctx) error {
	list := new(stats.List)
	list.Data = static.Ideologys

	data, err := proto.Marshal(list)
	if err != nil {
		return c.SendStatus(fasthttp.StatusConflict)
	}

	c.SendStatus(fasthttp.StatusOK)
	return c.Send(data)
}
