package repository

import "github.com/dnwandana/expense-tracker/model/entity"

type ExpenseRepository interface {
	Create(expense *entity.Expense)
	FindByID(expenseID int) *entity.Expense
	FindByUserID(userID int) []*entity.Expense
	Update(userID int, expense *entity.Expense)
	Delete(userID, expenseID int)
}
