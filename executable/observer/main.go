package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed public
var public embed.FS

func main() {
	app := fiber.New()

	app.Use(filesystem.New(filesystem.Config{
		Root: http.FS(public),
	}))

	if err := app.Listen(":10000"); err != nil {
		log.Fatal(err)
	}
}
