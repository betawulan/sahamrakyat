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

// Create ...
func (u userRepository) Create(ctx context.Context, user model.User) (model.User, error) {
	user.FirstOrder = time.Now()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.DeletedAt = time.Now()

	query, args, err := sq.Insert("user").
		Columns("full_name",
			"first_order",
			"created_at",
			"updated_at",
			"deleted_at").
		Values(user.FullName,
			user.FirstOrder,
			user.CreatedAt,
			user.UpdatedAt,
			user.DeletedAt).
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
