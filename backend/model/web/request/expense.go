package request

type ExpenseRequest struct {
	CategoryID  int    `json:"category_id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}
