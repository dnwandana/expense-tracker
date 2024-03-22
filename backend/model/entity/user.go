package entity

import "time"

// User entity is a struct to represent user data in the database
type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
