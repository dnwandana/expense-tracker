package request

type Expense struct {
	CategoryID  int     `json:"category_id"`
	Title       string  `json:"title"`
	Amount      int     `json:"amount"`
	Description *string `json:"description"`
}
