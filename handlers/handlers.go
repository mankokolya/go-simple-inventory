package handlers

import (
	"net/http"

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

func GetItemByID(c *fiber.Ctx) error {
	var itemID string = c.Params("id")

	item, err := services.GetItemById(itemID)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.Response[any]{
			Success: false,
			Message: err.Error(),
		})
	}

	return c.JSON(models.Response[models.Item]{
		Success: true,
		Message: "item found",
		Data:    item,
	})
}
