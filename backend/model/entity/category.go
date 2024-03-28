package entity

import "time"

// Category entity is a struct to represent category data in the database
type Category struct {
	ID        int
	UserID    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
