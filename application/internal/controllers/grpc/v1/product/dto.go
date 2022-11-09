package product

import pb_online_store_products "github.com/Meystergod/online-store-grpc-contracts/gen/go/online_store/products/v1"

type CreateProductDTO struct {
	Brand         string
	Title         string
	Description   string
	Price         string
	ImageId       *string
	Specification string
	CategoryId    uint32
	DiscountId    uint32
	InventoryId   uint32
}

func NewCreateProductDTOFromPB(product *pb_online_store_products.CreateProductRequest) *CreateProductDTO {
	return &CreateProductDTO{
		Brand:         product.GetBrand(),
		Title:         product.GetTitle(),
		Description:   product.GetDescription(),
		Price:         product.GetPrice(),
		ImageId:       product.ImageId,
		Specification: product.GetSpecifications(),
		CategoryId:    product.GetCategoryId(),
		DiscountId:    product.GetCategoryId(),
		InventoryId:   product.GetInventoryId(),
	}
}
