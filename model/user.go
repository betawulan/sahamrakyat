package model

import "time"

type User struct {
	ID         int64     `json:"id"`
	FullName   string    `json:"fullname"`
	FirstOrder string    `json:"first_order"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	DeletedAt  time.Time `json:"-"`
}

type UserFilter struct {
	Limit uint64
	Page  int
}

type UserResponse struct {
	Users []User `json:"users"`
}
