package policy

import (
	"context"

	"github.com/Meystergod/online-store-api/application/internal/controllers/grpc/v1/product"
	"github.com/Meystergod/online-store-api/application/internal/domain/product/model"
)

type productService interface {
	All(ctx context.Context) ([]model.Product, error)
	Create(ctx context.Context, dto product.CreateProductDTO) (model.Product, error)
}

type ProductPolicy struct {
	productService productService
}
