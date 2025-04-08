package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mankokolya/go-simple-inventory/routes"
)

func main() {
	var app *fiber.App = fiber.New()

	routes.SetupRoutes(app)

	app.Listen("127.0.0.1:8080")
}
