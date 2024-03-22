package entity

import "time"

type Expense struct {
	ID           string
	UserID       string
	CategoryID   string
	CategoryName string
	Amount       float64
	Description  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
