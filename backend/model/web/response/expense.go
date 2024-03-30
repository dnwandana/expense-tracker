package response

import (
	"time"
)

type Expense struct {
	ID          int       `json:"id"`
	CategoryID  int       `json:"category_id"`
	Category    string    `json:"category"`
	Title       string    `json:"title"`
	Amount      int       `json:"amount"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
