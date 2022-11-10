package sort

import (
	"strings"
)

type Order string

const (
	OrderASC  Order = "ASC"
	OrderDESC Order = "DESC"
)

type Sortable interface {
	Field() string
	Order() string
}

type Opts struct {
	field string
	order string
}

func NewOptions(field string) *Opts {
	sortOrder := OrderASC
	if strings.HasPrefix(field, "-") {
		sortOrder = OrderDESC
		field = strings.TrimPrefix(field, "-")
	}
	return &Opts{
		field: field,
		order: string(sortOrder),
	}
}

func (o *Opts) Field() string {
	return o.field
}
func (o *Opts) Order() string {
	return o.order
}
