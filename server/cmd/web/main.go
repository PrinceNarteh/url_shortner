package main

import (
	"url_shortener/pkg/handlers"
	"url_shortener/pkg/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	model.Setup()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/urls", handlers.GetAllRedirects)

	app.Listen(":4000")
}
