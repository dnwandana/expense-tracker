package repository

import (
	"database/sql"

	"github.com/dnwandana/expense-tracker/model/entity"
	"github.com/dnwandana/expense-tracker/model/web/response"
	"github.com/dnwandana/expense-tracker/utils"
)

type ExpenseRepositoryImpl struct {
	DB *sql.DB
}

func NewExpenseRepository(db *sql.DB) ExpenseRepository {
	return &ExpenseRepositoryImpl{DB: db}
}

func (repo *ExpenseRepositoryImpl) Create(expense *entity.Expense) {
	_, err := repo.DB.Exec("INSERT INTO expenses (user_id, category_id, title, amount, description) VALUES (?, ?, ?, ?, ?)", expense.UserID, expense.CategoryID, expense.Title, expense.Amount, expense.Description)
	utils.PanicIfError(err)
}

func (repo *ExpenseRepositoryImpl) FindByID(expenseID int) *response.Expense {
	query := "SELECT expenses.id, expenses.category_id, categories.name, expenses.title, expenses.amount, expenses.description, expenses.created_at, expenses.updated_at FROM expenses INNER JOIN categories ON expenses.category_id = categories.id WHERE expenses.id = ?"
	row, err := repo.DB.Query(query, expenseID)
	utils.PanicIfError(err)
	defer row.Close()

	expense := new(response.Expense)
	if row.Next() {
		var description sql.NullString
		err = row.Scan(&expense.ID, &expense.CategoryID, &expense.Category, &expense.Title, &expense.Amount, &description, &expense.CreatedAt, &expense.UpdatedAt)

		if description.Valid {
			expense.Description = description.String
		}

		utils.PanicIfError(err)
	}

	return expense
}

func (repo *ExpenseRepositoryImpl) FindByUserID(userID int) []*response.Expense {
	query := "SELECT expenses.id, expenses.category_id, categories.name, expenses.title, expenses.amount, expenses.description, expenses.created_at, expenses.updated_at FROM expenses INNER JOIN categories ON expenses.category_id = categories.id WHERE expenses.user_id = ?"
	rows, err := repo.DB.Query(query, userID)
	utils.PanicIfError(err)
	defer rows.Close()

	expenses := make([]*response.Expense, 0)
	for rows.Next() {
		expense := new(response.Expense)
		var description sql.NullString

		err = rows.Scan(&expense.ID, &expense.CategoryID, &expense.Category, &expense.Title, &expense.Amount, &description, &expense.CreatedAt, &expense.UpdatedAt)

		if description.Valid {
			expense.Description = description.String
		}

		utils.PanicIfError(err)
		expenses = append(expenses, expense)
	}

	return expenses
}

func (repo *ExpenseRepositoryImpl) Update(userID int, expense *entity.Expense) {
	_, err := repo.DB.Exec("UPDATE expenses SET category_id = ?, title = ?, amount = ?, description = ? WHERE user_id = ? AND id = ?", expense.CategoryID, expense.Title, expense.Amount, expense.Description, userID, expense.ID)
	utils.PanicIfError(err)
}

func (repo *ExpenseRepositoryImpl) Delete(userID, expenseID int) {
	_, err := repo.DB.Exec("DELETE FROM expenses WHERE user_id = ? AND id = ?", userID, expenseID)
	utils.PanicIfError(err)
}
