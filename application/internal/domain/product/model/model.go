package model

import "time"

type Product struct {
	Id             uint32     `json:"id"`
	Brand          string     `json:"brand"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	Price          string     `json:"price"`
	ImageId        *string    `json:"image_id"`
	Specifications string     `json:"specifications"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	CategoryId     uint32     `json:"category_id"`
	DiscountId     uint32     `json:"discount_id"`
	InventoryId    uint32     `json:"inventory_id"`
}
