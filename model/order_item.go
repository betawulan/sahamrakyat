package model

import "time"

type OrderItem struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	ExpiredAt string    `json:"expired_at"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

type OrderItemFilter struct {
	Limit  uint64
	Page   int
	Status []int
}

type OrderItemResponse struct {
	OrdersItem []OrderItem `json:"orders_item"`
}
