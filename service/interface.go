package service

import (
	"context"
	"github.com/betawulan/sahamrakyat/model"
)

type UserService interface {
	Create(ctx context.Context, user model.User) (model.User, error)
	ReadByID(ctx context.Context, ID int64) (model.User, error)
	Update(ctx context.Context, ID int64, user model.User) error
	Delete(ctx context.Context, ID int64) error
}
