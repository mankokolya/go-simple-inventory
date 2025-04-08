package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mankokolya/go-simple-inventory/models"
)

var storage []models.Item = []models.Item{}

func GetAllItems() []models.Item {
	return storage
}

func GetItemById(id string) (models.Item, error) {
	for _, item := range storage {
		if item.ID == id {
			return item, nil
		}
	}

	return models.Item{}, errors.New("item not found")
}

func CreateItem(itemRequest models.ItemRequest) models.Item {
	var newItem models.Item = models.Item{
		ID:        uuid.New().String(),
		Name:      itemRequest.Name,
		Price:     itemRequest.Price,
		Quantity:  itemRequest.Quantity,
		CreatedAt: time.Now(),
	}

	storage = append(storage, newItem)

	return newItem
}

func UpdateItem(itemRequest models.ItemRequest, id string) (models.Item, error) {
	for index, item := range storage {
		if item.ID == id {
			item.Name = itemRequest.Name
			item.Price = itemRequest.Price
			item.Quantity = itemRequest.Quantity
			item.UpdatedAt = time.Now()

			storage[index] = item

			return item, nil
		}
	}

	return models.Item{}, errors.New("item update failed, item not found")
}

func DeleteItem(id string) bool {
	var newItems []models.Item = []models.Item{}

	var itemPresent bool = false

	for _, item := range storage {
		if item.ID != id {
			newItems = append(newItems, item)
		} else {
			itemPresent = true
		}
	}

	if itemPresent {
		storage = newItems
	}

	return itemPresent
}
