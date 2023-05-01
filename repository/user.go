package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/betawulan/sahamrakyat/model"
)

type userRepository struct {
	db *sql.DB
}

func (u userRepository) ReadByID(ctx context.Context, ID int64) (model.User, error) {
	query, args, err := sq.Select("id",
		"fullname",
		"first_order").
		From("user").
		Where(sq.Eq{"id": ID}).
		ToSql()
	if err != nil {
		return model.User{}, err
	}

	row := u.db.QueryRowContext(ctx, query, args...)
	var user model.User
	err = row.Scan(&user.ID,
		&user.FullName,
		&user.FirstOrder)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u userRepository) Update(ctx context.Context, ID int64, user model.User) error {
	query, args, err := sq.Update("user").
		Set("fullname", user.FullName).
		Set("first_order", user.FirstOrder).
		Set("updated_at", time.Now()).
		Where(sq.Eq{"id": ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = u.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (u userRepository) Delete(ctx context.Context, ID int64) error {
	query, args, err := sq.Update("user").
		Set("deleted_at", time.Now()).
		Set("status_deleted", true).
		Where(sq.Eq{"id": ID}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = u.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Create ...
func (u userRepository) Create(ctx context.Context, user model.User) (model.User, error) {
	user.CreatedAt = time.Now()

	query, args, err := sq.Insert("user").
		Columns("fullname",
			"first_order",
			"created_at").
		Values(user.FullName,
			user.FirstOrder,
			user.CreatedAt).
		ToSql()
	if err != nil {
		return model.User{}, err
	}

	res, err := u.db.ExecContext(ctx, query, args...)
	if err != nil {
		return model.User{}, err
	}

	user.ID, err = res.LastInsertId()
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u userRepository) Read(ctx context.Context, filter model.UserFilter) ([]model.User, error) {
	querySelect := sq.Select("id",
		"fullname",
		"first_order").
		From("user")

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

	rows, err := u.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := make([]model.User, 0)
	for rows.Next() {
		var user model.User

		err = rows.Scan(
			&user.ID,
			&user.FullName,
			&user.FirstOrder)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func NewUserRepository(db *sql.DB) UserRepository {
	return userRepository{
		db: db,
	}
}
