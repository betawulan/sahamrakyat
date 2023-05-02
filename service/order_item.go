package service

import (
	"context"

	"github.com/betawulan/sahamrakyat/model"
	"github.com/betawulan/sahamrakyat/repository"
)

type orderItemService struct {
	orderItemRepo repository.OrderItemRepository
}

func (o orderItemService) Publish(ctx context.Context, ID int64) error {
	err := o.orderItemRepo.Publish(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (o orderItemService) UnPublish(ctx context.Context, ID int64) error {
	err := o.orderItemRepo.UnPublish(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (o orderItemService) Read(ctx context.Context, filter model.OrderItemFilter) (model.OrderItemResponse, error) {
	ordersItem, err := o.orderItemRepo.Read(ctx, filter)
	if err != nil {
		return model.OrderItemResponse{}, err
	}

	return model.OrderItemResponse{
		OrdersItem: ordersItem,
	}, nil
}

func (o orderItemService) Create(ctx context.Context, orderItem model.OrderItem) (model.OrderItem, error) {
	orderItem, err := o.orderItemRepo.Create(ctx, orderItem)
	if err != nil {
		return model.OrderItem{}, err
	}

	return orderItem, nil
}

func (o orderItemService) ReadByID(ctx context.Context, ID int64) (model.OrderItem, error) {
	orderItem, err := o.orderItemRepo.ReadByID(ctx, ID)
	if err != nil {
		return model.OrderItem{}, err
	}

	return orderItem, nil
}

func (o orderItemService) Update(ctx context.Context, ID int64, orderItem model.OrderItem) error {
	err := o.orderItemRepo.Update(ctx, ID, orderItem)
	if err != nil {
		return err
	}

	return nil
}

func NewOrderItemService(orderItemRepository repository.OrderItemRepository) OrderItemService {
	return orderItemService{
		orderItemRepo: orderItemRepository,
	}
}
