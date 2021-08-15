package main

import (
	"alcor/service/security"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func main() {
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

	app.Listen(":9999")
}
