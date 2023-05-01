package service

import (
	"context"

	"github.com/betawulan/sahamrakyat/model"
	"github.com/betawulan/sahamrakyat/repository"
)

type orderItemService struct {
	orderItemRepo repository.OrderItemRepository
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

func (o orderItemService) Delete(ctx context.Context, ID int64) error {
	//TODO implement me
	panic("implement me")
}

func NewOrderItemService(orderItemRepository repository.OrderItemRepository) OrderItemService {
	return orderItemService{
		orderItemRepo: orderItemRepository,
	}
}
