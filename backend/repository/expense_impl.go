package repository

import (
	"database/sql"

	"github.com/dnwandana/expense-tracker/model/entity"
	"github.com/dnwandana/expense-tracker/utils"
)

type ExpenseRepositoryImpl struct {
	DB *sql.DB
}

func NewExpenseRepository(db *sql.DB) ExpenseRepository {
	return &ExpenseRepositoryImpl{DB: db}
}

func (repo *ExpenseRepositoryImpl) Create(expense *entity.Expense) {
	_, err := repo.DB.Exec("INSERT INTO expenses (id, user_id, category_id, amount, description) VALUES (?, ?, ?, ?, ?)", expense.ID, expense.UserID, expense.CategoryID, expense.Amount, expense.Description)
	utils.PanicIfError(err)
}

func (repo *ExpenseRepositoryImpl) FindByID(expenseID int) *entity.Expense {
	query := "SELECT id, category_id, amount, description, created_at, updated_at FROM expenses WHERE id = ?"
	row, err := repo.DB.Query(query, expenseID)
	utils.PanicIfError(err)
	defer row.Close()

	expense := new(entity.Expense)
	if row.Next() {
		err = row.Scan(&expense.ID, &expense.CategoryID, &expense.Amount, &expense.Description, &expense.CreatedAt, &expense.UpdatedAt)
		utils.PanicIfError(err)
	}

	return expense
}

func (repo *ExpenseRepositoryImpl) FindByUserID(userID int) []*entity.Expense {
	query := "SELECT id, category_id, amount, description, created_at, updated_at FROM expenses WHERE user_id = ?"
	rows, err := repo.DB.Query(query, userID)
	utils.PanicIfError(err)
	defer rows.Close()

	expenses := make([]*entity.Expense, 0)
	for rows.Next() {
		expense := new(entity.Expense)
		err = rows.Scan(&expense.ID, &expense.CategoryID, &expense.Amount, &expense.Description, &expense.CreatedAt, &expense.UpdatedAt)
		utils.PanicIfError(err)
		expenses = append(expenses, expense)
	}

	return expenses
}

func (repo *ExpenseRepositoryImpl) Update(userID int, expense *entity.Expense) {
	_, err := repo.DB.Exec("UPDATE expenses SET category_id = ?, amount = ?, description = ? WHERE user_id = ? AND id = ?", expense.CategoryID, expense.Amount, expense.Description, userID, expense.ID)
	utils.PanicIfError(err)
}

func (repo *ExpenseRepositoryImpl) Delete(userID, expenseID int) {
	_, err := repo.DB.Exec("DELETE FROM expenses WHERE user_id = ? AND id = ?", userID, expenseID)
	utils.PanicIfError(err)
}
