package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/betawulan/sahamrakyat/model"
)

type orderItemRepository struct {
	db *sql.DB
}

func (o orderItemRepository) Create(ctx context.Context, orderItem model.OrderItem) (model.OrderItem, error) {
	orderItem.CreatedAt = time.Now()

	query, args, err := sq.Insert("orders_item").
		Columns("name",
			"price",
			"expired_at",
			"created_at").
		Values(orderItem.Name,
			orderItem.Price,
			orderItem.ExpiredAt,
			orderItem.CreatedAt).
		ToSql()
	if err != nil {
		return model.OrderItem{}, err
	}

	res, err := o.db.ExecContext(ctx, query, args...)
	if err != nil {
		return model.OrderItem{}, err
	}

	orderItem.ID, err = res.LastInsertId()
	if err != nil {
		return model.OrderItem{}, err
	}

	return orderItem, nil
}

func (o orderItemRepository) ReadByID(ctx context.Context, ID int64) (model.OrderItem, error) {
	query, args, err := sq.Select("id",
		"name",
		"price",
		"expired_at").
		From("orders_item").
		Where(sq.Eq{"id": ID}).
		ToSql()
	if err != nil {
		return model.OrderItem{}, err
	}

	row := o.db.QueryRowContext(ctx, query, args...)
	var orderItem model.OrderItem
	err = row.Scan(&orderItem.ID,
		&orderItem.Name,
		&orderItem.Price,
		&orderItem.ExpiredAt)
	if err != nil {
		return model.OrderItem{}, err
	}

	return orderItem, nil
}

func (o orderItemRepository) Update(ctx context.Context, ID int64, orderItem model.OrderItem) error {
	//TODO implement me
	panic("implement me")
}

func (o orderItemRepository) Delete(ctx context.Context, ID int64) error {
	//TODO implement me
	panic("implement me")
}

func NewOrderItemRepository(db *sql.DB) OrderItemRepository {
	return orderItemRepository{
		db: db,
	}
}
