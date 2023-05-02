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

func (o orderItemRepository) Publish(ctx context.Context, ID int64) error {
	query, args, err := sq.Update("orders_item").
		Set("deleted_at", time.Now()).
		Set("status_deleted", true).
		Where(sq.Eq{"id": ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = o.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (o orderItemRepository) UnPublish(ctx context.Context, ID int64) error {
	query, args, err := sq.Update("orders_item").
		Set("deleted_at", time.Now()).
		Set("status_deleted", false).
		Where(sq.Eq{"id": ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = o.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (o orderItemRepository) Read(ctx context.Context, filter model.OrderItemFilter) ([]model.OrderItem, error) {
	querySelect := sq.Select("id",
		"name",
		"price",
		"expired_at").
		From("orders_item")

	if filter.Limit != 0 {
		querySelect = querySelect.Limit(filter.Limit)
	}

	if filter.Page != 0 {
		querySelect = querySelect.Offset((uint64(filter.Page) - 1) * filter.Limit)
	}

	query, args, err := querySelect.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := o.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	ordersItem := make([]model.OrderItem, 0)
	for rows.Next() {
		var orderItem model.OrderItem

		err = rows.Scan(
			&orderItem.ID,
			&orderItem.Name,
			&orderItem.Price,
			&orderItem.ExpiredAt)
		if err != nil {
			return nil, err
		}

		ordersItem = append(ordersItem, orderItem)
	}

	return ordersItem, nil
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
	query, args, err := sq.Update("orders_item").
		Set("name", orderItem.Name).
		Set("price", orderItem.Price).
		Set("expired_at", orderItem.ExpiredAt).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = o.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func NewOrderItemRepository(db *sql.DB) OrderItemRepository {
	return orderItemRepository{
		db: db,
	}
}
