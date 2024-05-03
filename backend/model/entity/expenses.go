package entity

import (
	"database/sql"
	"time"
)

// Expense entity is a struct to represent expense data in the database
type Expense struct {
	ID          int
	UserID      int
	CategoryID  int
	Title       string
	Amount      int
	Description sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
