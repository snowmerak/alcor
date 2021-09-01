package main

import (
	"alcor/db"
	"alcor/service/admin"
	"alcor/service/bundle"
	"alcor/service/candidate"
	paper_info "alcor/service/paper/info"
	"alcor/service/paper/vote"
	"alcor/service/security"
	"alcor/service/survey"
	voter_info "alcor/service/voter/info"
	"alcor/service/voter/register"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
	db.Init()

	app := fiber.New()

	adminPassword := os.Getenv("ADMIN_PASSWORD")

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
		AllowMethods:     "GET, POST, DELETE, PATCH",
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

	app.Get("/bundle", bundle.GetLast)
	app.Get("/bundle/:id", bundle.Get)

	app.Post("/survey/reply", survey.PostService)

	app.Get("/survey/regions", survey.GetRegions)
	app.Get("/survey/jobs", survey.GetJobs)
	app.Get("/survey/educations", survey.GetEducations)
	app.Get("/survey/house_types", survey.GetsHouseTypes)
	app.Get("/survey/ideologys", survey.GetIdeologys)

	app.Get("/survey/last_number", survey.GetLastNumber)
	app.Get("/survey/:number", survey.GetIndex)

	app.Post("/candidate", admin.MakePostCandidate(adminPassword))
	app.Patch("/candidate", admin.MakePatchCandidate(adminPassword))
	app.Delete("/candidate", admin.MakeDeleteCandidate(adminPassword))
	app.Post("/admin", admin.MakeCheckPassword(adminPassword))

	app.Listen(":9999")
}
