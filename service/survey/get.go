package survey

import (
	"alcor/db"
	"alcor/model/stats"
	"alcor/static"
	"alcor/worker/rdb"
	"context"
	"strconv"

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

func GetMaxCount(c *fiber.Ctx) error {
	data, err := rdb.GetSurveyIndex()
	if err != nil {
		data = 0
	}
	c.SendStatus(fasthttp.StatusOK)
	return c.SendString(strconv.FormatUint(data, 10))
}

func GetIndex(c *fiber.Ctx) error {
	n, err := strconv.ParseUint(c.Params("number"), 10, 64)
	if err != nil {
		return c.SendStatus(fasthttp.StatusBadRequest)
	}
	data := new(db.Data)
	data.Number = int64(n)
	if err := db.SelectData(context.Background(), data); err != nil {
		return c.SendStatus(fasthttp.StatusBadRequest)
	}
	survey := new(stats.Data)
	survey.Age = int32(data.Age)
	survey.Candidate = data.Candidate
	survey.Divorced = data.Divorced
	survey.Education = data.Education
	survey.Gender = data.Gender
	survey.HasCar = data.HasCar
	survey.HasDebt = data.HasDebt
	survey.HouseType = data.HouseType
	survey.Ideology = data.Ideology
	survey.Job = data.Job
	survey.Married = data.Married
	survey.Region = data.Region
	survey.Salary = int32(data.Salary)
	bs, err := proto.Marshal(survey)
	if err != nil {
		return c.SendStatus(fasthttp.StatusConflict)
	}
	c.SendStatus(fasthttp.StatusOK)
	return c.Send(bs)
}
