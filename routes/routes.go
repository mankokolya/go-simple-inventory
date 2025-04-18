package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mankokolya/go-simple-inventory/handlers"
	"github.com/mankokolya/go-simple-inventory/middlewares"
)

func SetupRoutes(app *fiber.App) {

	//public routes
	app.Post("/api/v1/signup", handlers.Signup)
	app.Post("/api/v1/login", handlers.Login)
	app.Get("/api/v1/items", handlers.GetAllItems)
	app.Get("/api/v1/items/:id", handlers.GetItemByID)

	//private routes authentication is required
	app.Post("/api/v1/items", middlewares.CreateMiddleware(), handlers.CreateItem)
	app.Put("/api/v1/items/:id", middlewares.CreateMiddleware(), handlers.UpdateItem)
	app.Delete("/api/v1/items/:id", middlewares.CreateMiddleware(), handlers.DeleteItem)
}
