package storage

import (
	"context"

	"github.com/Meystergod/online-store-api/application/internal/domain/product/model"
	db "github.com/Meystergod/online-store-api/application/pkg/client/postgresql/model"
	"github.com/Meystergod/online-store-api/application/pkg/logging"

	sq "github.com/Masterminds/squirrel"
)

type ProductStorage struct {
	queryBuilder sq.StatementBuilderType
	pgClient     PostgreSQLClient
}

func NewProductStorage(client PostgreSQLClient) ProductStorage {
	return ProductStorage{
		queryBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		pgClient:     client,
	}
}

const (
	scheme = "public"
	table  = "product"
)

func (s *ProductStorage) All(ctx context.Context) ([]model.Product, error) {
	query := s.queryBuilder.Select("id").
		Columns("brand").
		Columns("title").
		Columns("description").
		Columns("price").
		Columns("image_id").
		Columns("specifications").
		Columns("created_at").
		Columns("updated_at").
		Columns("category_id").
		Columns("discount_id").
		Columns("inventory_id").
		From(scheme + "." + table)

	// will be filtering and sorting

	sql, args, err := query.ToSql()
	logger := logging.GetLogger(ctx).WithFields(map[string]interface{}{
		"sql":   sql,
		"table": table,
		"args":  args,
	})

	logger.Trace("do query")

	rows, err := s.pgClient.Query(ctx, sql, args...)
	if err != nil {
		err = db.ErrorDoQuery(err)
		logger.Error(err)
		return nil, err
	}

	defer rows.Close()

	list := make([]model.Product, 0)

	for rows.Next() {
		p := model.Product{}
		if err = rows.Scan(
			&p.Id,
			&p.Brand,
			&p.Title,
			&p.Description,
			&p.Price,
			&p.ImageId,
			&p.Specifications,
			&p.CreatedAt,
			&p.UpdatedAt,
			&p.CategoryId,
			&p.DiscountId,
			&p.InventoryId,
		); err != nil {
			err = db.ErrorScan(err)
			logger.Error(err)
			return nil, err
		}

		list = append(list, p)
	}
	return list, nil
}
