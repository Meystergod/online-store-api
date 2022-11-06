package model

import "time"

type Product struct {
	Id             int        `json:"id"`
	Brand          string     `json:"brand"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	Price          int        `json:"price"`
	ImageId        *string    `json:"image_id"`
	Specifications string     `json:"specifications"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	CategoryId     int        `json:"category_id"`
	DiscountId     int        `json:"discount_id"`
	InventoryId    int        `json:"inventory_id"`
}
