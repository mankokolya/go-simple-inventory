package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	var app *fiber.App = fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world!")
	})

	app.Listen("127.0.0.1:8080")
}
