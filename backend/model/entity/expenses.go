package entity

import "time"

// Expense entity is a struct to represent expense data in the database
type Expense struct {
	ID           int
	UserID       int
	CategoryID   int
	CategoryName string
	Amount       int
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
