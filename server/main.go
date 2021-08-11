package main

import (
	"alcor/db"
	"alcor/service/candidate"
	"alcor/service/enroll"
	"alcor/service/register"
	"alcor/service/vote"
	"alcor/worker/hashbuffer"
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/snowmerak/jumper"
	"google.golang.org/protobuf/proto"
)

func main() {
	max, err := strconv.Atoi(os.Getenv("MAX_CONNECTION"))
	if err != nil {
		max = 100000
	}
	expiration, err := strconv.Atoi(os.Getenv("EXPIRATION_TIME"))
	if err != nil {
		expiration = 5 * int(time.Minute)
	}
	expiration *= int(time.Second)

	certLogger, err := jumper.New("cert_error", enroll.NewCertError(""))
	if err != nil {
		log.Fatal(err)
	}
	go certLogger.Start(10000)
	defer certLogger.Close()

	hashBufferLogger, err := jumper.New("hashBuffer_error", hashbuffer.NewRingError(""))
	if err != nil {
		log.Fatal(err)
	}
	go hashBufferLogger.Start(10000)
	defer hashBufferLogger.Close()

	registerLogger, err := jumper.New("reg_error", register.NewRegisterError(""))
	if err != nil {
		log.Fatal(err)
	}
	go registerLogger.Start(10000)
	defer registerLogger.Close()

	app := fiber.New()

	//admin server
	go func() {
		admin := fiber.New()
		admin.Post("/register", func(c *fiber.Ctx) error {
			data := new(candidate.Candidate)
			if err := proto.Unmarshal(c.Body(), data); err != nil {
				return c.SendString(err.Error())
			}
			candidate := new(db.Candidate)
			candidate.Name = data.Name
			candidate.Markdown = data.Markdown
			return c.SendString(db.InsertCandidate(context.Background(), candidate).Error())
		})
		admin.Get("/stop", func(c *fiber.Ctx) error {
			return app.Shutdown()
		})
		admin.Get("/count")
		if err := admin.Listen(":10000"); err != nil {
			log.Fatal(err)
		}
	}()

	app.Use(limiter.New(limiter.Config{
		Max:        max,
		Expiration: time.Duration(expiration),
	}), cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowMethods:     "GET",
	}))

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello!")
	})

	app.Get("/enroll/phone", enroll.PhoneCertSocket)
	app.Get("/vote", vote.PhoneCertSocket)
	app.Get("/register", register.PhoneRegisterSocket)
	app.Get("/candidate", candidate.GetList)
	app.Get("/candidate/:name", candidate.Get)

	db.Init()

	if err := app.Listen(":9999"); err != nil {
		log.Fatal(err)
	}
}
