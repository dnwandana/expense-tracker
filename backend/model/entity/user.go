package entity

import "time"

type User struct {
	ID        string
	Fullname  string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
