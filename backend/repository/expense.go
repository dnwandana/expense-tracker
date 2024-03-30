package repository

import (
	"github.com/dnwandana/expense-tracker/model/entity"
	"github.com/dnwandana/expense-tracker/model/web/response"
)

type ExpenseRepository interface {
	Create(expense *entity.Expense)
	FindByID(expenseID int) *response.Expense
	FindByUserID(userID int) []*response.Expense
	Update(userID int, expense *entity.Expense)
	Delete(userID, expenseID int)
}
