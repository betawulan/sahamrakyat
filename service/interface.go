package service

import (
	"context"
	"github.com/betawulan/sahamrakyat/model"
)

type UserService interface {
	Create(ctx context.Context, user model.User) (model.User, error)
	ReadByID(ctx context.Context, ID int64) (model.User, error)
	Update(ctx context.Context, ID int64, user model.User) error
	Read(ctx context.Context, filter model.UserFilter) (model.UserResponse, error)
	Publish(ctx context.Context, ID int64) error   // soft delete
	UnPublish(ctx context.Context, ID int64) error // soft delete
}

type OrderItemService interface {
	Create(ctx context.Context, orderItem model.OrderItem) (model.OrderItem, error)
	ReadByID(ctx context.Context, ID int64) (model.OrderItem, error)
	Update(ctx context.Context, ID int64, orderItem model.OrderItem) error
	Read(ctx context.Context, filter model.OrderItemFilter) (model.OrderItemResponse, error)
	Publish(ctx context.Context, ID int64) error   // soft delete
	UnPublish(ctx context.Context, ID int64) error // soft delete
}
