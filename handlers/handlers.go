package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mankokolya/go-simple-inventory/models"
	"github.com/mankokolya/go-simple-inventory/services"
)

func GetAllItems(c *fiber.Ctx) error {
	var items []models.Item = services.GetAllItems()

	return c.JSON(models.Response[[]models.Item]{
		Success: true,
		Message: "All items data",
		Data:    items,
	})
}
