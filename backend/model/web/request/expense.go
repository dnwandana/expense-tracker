package request

type ExpenseRequest struct {
	CategoryID  string  `json:"category_id"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}
