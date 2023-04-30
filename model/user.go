package model

import "time"

type User struct {
	ID         uint64    `json:"id"`
	FullName   string    `json:"fullName"`
	FirstOrder time.Time `json:"-"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	DeletedAt  time.Time `json:"-"`
}
