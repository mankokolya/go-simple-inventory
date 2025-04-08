package models

import "time"

type Item struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
