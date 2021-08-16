package stats

import (
	"alcor/db"
	"alcor/model/stats"
	"alcor/service/security"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"google.golang.org/protobuf/proto"
)

func StatsService(c *fiber.Ctx) error {
	_, data, err := security.Decapsulate(c.Body())
	if err != nil {
		log.Println(err)
		return c.SendStatus(fasthttp.StatusBadRequest)
	}

	stats := new(stats.Data)
	if err := proto.Unmarshal(data, stats); err != nil {
		log.Println(err)
		return c.SendStatus(fasthttp.StatusBadRequest)
	}

	voter := new(db.Voter)
	voter.HashID = stats.VoterID
	if err := db.SelectVoter(context.Background(), voter); err != nil {
		log.Println(err)
		return c.SendStatus(fasthttp.StatusConflict)
	}

	if voter.RepliedSurvey {
		log.Println(err)
		return c.SendStatus(fasthttp.StatusBadRequest)
	}

	statsdb := new(db.Data)
	statsdb.Age = int16(stats.Age)
	statsdb.Candidate = stats.Candidate
	statsdb.Divorced = stats.Divorced
	statsdb.Education = int16(stats.Education)
	statsdb.Gender = stats.Gender
	statsdb.HasCar = stats.HasCar
	statsdb.HasDebt = stats.HasDebt
	statsdb.HouseType = int16(stats.HouseType)
	statsdb.Ideology = int16(stats.Ideology)
	statsdb.Job = int16(stats.Job)
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
