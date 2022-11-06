package model

type Product struct {
	Id             int    `json:"id"`
	Brand          string `json:"brand"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Price          int    `json:"price"`
	ImageId        int    `json:"image_id"`
	Specifications string `json:"specifications"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	CategoryId     int    `json:"category_id"`
	DiscountId     int    `json:"discount_id"`
	InventoryId    int    `json:"inventory_id"`
}
