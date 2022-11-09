package storage

import (
	"database/sql"
	"time"
)

type Product struct {
	Id             uint32
	Brand          string
	Title          string
	Description    string
	Price          string
	ImageId        sql.NullString
	Specifications string
	CreatedAt      time.Time
	UpdatedAt      *time.Time
	CategoryId     uint32
	DiscountId     uint32
	InventoryId    uint32
}

type CreateProductDTO struct {
	Brand          string
	Title          string
	Description    string
	Price          string
	ImageId        sql.NullString
	Specifications string
	CreatedAt      time.Time
	UpdatedAt      *time.Time
	CategoryId     uint32
	DiscountId     uint32
	InventoryId    uint32
}
