package model

import "time"

type User struct {
	ID         int64     `json:"id"`
	FullName   string    `json:"fullName"`
	FirstOrder string    `json:"firstOrder"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	DeletedAt  time.Time `json:"-"`
}
