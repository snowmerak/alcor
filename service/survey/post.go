package survey

import (
	"alcor/db"
	"alcor/model/stats"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/proto"
)

func PostService(c *fiber.Ctx) error {
	stats := new(stats.Data)
	if err := proto.Unmarshal(c.Body(), stats); err != nil {
		log.Println(err)
		return c.SendStatus(fasthttp.StatusBadRequest)
	}

	voter := new(db.Voter)
	voter.HashID = stats.VoterID
	if err := db.SelectVoter(context.Background(), voter); err != nil {
		log.Println(err)
		return c.SendStatus(fasthttp.StatusNotFound)
	}

	if voter.RepliedSurvey {
		log.Println("already submitted")
		return c.SendStatus(fasthttp.StatusBadRequest)
	}

	statsdb := new(db.Data)
	statsdb.Age = int16(stats.Age)
	statsdb.Region = stats.Region
	statsdb.Salary = int16(stats.Salary)
	statsdb.Candidate = stats.Candidate
	statsdb.Divorced = stats.Divorced
	statsdb.Education = stats.Education
	statsdb.Gender = stats.Gender
	statsdb.HasCar = stats.HasCar
	statsdb.HasDebt = stats.HasDebt
	statsdb.HouseType = stats.HouseType
	statsdb.Ideology = stats.Ideology
	statsdb.Job = stats.Job
	statsdb.Married = stats.Married
	if err := db.InsertData(context.Background(), statsdb); err != nil {
		log.Println(err)
		return c.SendStatus(fasthttp.StatusConflict)
	}

	if err := db.UpdateVoterRepliedSurvey(context.Background(), voter); err != nil {
		log.Println(err)
	}

	return c.SendStatus(fasthttp.StatusOK)
}
