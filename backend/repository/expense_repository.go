package repository

import "github.com/dnwandana/expense-tracker/model/entity"

type ExpenseRepository interface {
	Create(expense *entity.Expense)
	FindByID(expenseID string) *entity.Expense
	FindByUserID(userID string) []*entity.Expense
	Update(userID string, expense *entity.Expense)
	Delete(userID, expenseID string)
}
