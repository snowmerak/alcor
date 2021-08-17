package main

import (
	"alcor/db"
	"alcor/service/candidate"
	paper_info "alcor/service/paper/info"
	"alcor/service/paper/vote"
	"alcor/service/security"
	"alcor/service/survey"
	voter_info "alcor/service/voter/info"
	"alcor/service/voter/register"
	"context"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	db.Init()

	// test code
	db.InsertCandidate(context.Background(), &db.Candidate{
		Name:     "이슬기",
		Markdown: "## 나",
	})

	app := fiber.New()

	max, err := strconv.Atoi(os.Getenv("MAX_CONNECTION"))
	if err != nil {
		max = 100000
	}
	expiration, err := strconv.Atoi(os.Getenv("EXPIRATION_TIME"))
	if err != nil {
		expiration = 5 * int(time.Minute)
	}
	expiration *= int(time.Second)

	app.Use(limiter.New(limiter.Config{
		Max:        max,
		Expiration: time.Duration(expiration),
	}), cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST",
	}))

	app.Post("/handshake", security.Service)

	app.Post("/voter/register/request", register.PhoneService)
	app.Post("/voter/register/cert", register.CertService)

	app.Get("/voter/info/:id", voter_info.GetService)

	app.Get("/candidate", candidate.GetService)
	app.Get("/candidate/:name", candidate.GetAService)

	app.Post("/vote/request", vote.RequestService)
	app.Post("/vote/submit", vote.SubmitService)

	app.Get("/paper/info/:id", paper_info.GetService)

	app.Post("/survey/reply", survey.PostService)

	app.Listen(":9999")
}
