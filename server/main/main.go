package main

import (
	"mserv/routes/index"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api", func(c *fiber.Ctx) error {
		return c.Next()
	})

	api.Post("/new", index.NewPostHandler)

	app.Listen(":8082")
}
