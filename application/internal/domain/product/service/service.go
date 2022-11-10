package service

import (
	"context"
	"github.com/Meystergod/online-store-api/application/internal/controllers/grpc/v1/product"
	"github.com/Meystergod/online-store-api/application/internal/domain/product/model"

	"github.com/Meystergod/online-store-api/application/internal/domain/product/storage"
)

type repository interface {
	All(ctx context.Context) ([]storage.Product, error)
	Create(ctx context.Context, dto storage.CreateProductDTO) (storage.Product, error)
}

type Service struct {
	repository repository
}

func (s *Service) All(ctx context.Context) ([]model.Product, error) {

}

func (s *Service) Create(ctx context.Context, dto product.CreateProductDTO) (model.Product, error) {

}
