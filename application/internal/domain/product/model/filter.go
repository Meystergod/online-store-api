package model

import (
	"github.com/Meystergod/online-store-api/application/pkg/api/filter"
	"github.com/Meystergod/online-store-api/application/pkg/api/sort"
	pb_online_store_products "github.com/Meystergod/online-store-grpc-contracts/gen/go/online_store/products/v1"
)

const (
	paginationFilterField  = "pagination"
	brandFilterField       = "brand"
	titleFilterField       = "title"
	priceFilterField       = "price"
	categoryIdFilterField  = "category_id"
	discountIdFilterField  = "discount_id"
	inventoryIdFilterField = "inventory_id"
)

func productsFilterFields() map[string]string {
	return map[string]string{
		paginationFilterField:  filter.DataTypeStr,
		brandFilterField:       filter.DataTypeStr,
		titleFilterField:       filter.DataTypeStr,
		priceFilterField:       filter.DataTypeStr,
		categoryIdFilterField:  filter.DataTypeInt,
		discountIdFilterField:  filter.DataTypeInt,
		inventoryIdFilterField: filter.DataTypeInt,
	}
}

func ProductsSort(req *pb_online_store_products.GetAllProductRequest) sort.Sortable {
	field := req.GetSort().GetField()
	return sort.NewOptions(field)
}

func ProductsFilter(req *pb_online_store_products.GetAllProductRequest) filter.Filterable {
	return filter.NewOptions(req.GetPagination().GetLimit(), req.GetPagination().GetLimit(), productsFilterFields())
}
