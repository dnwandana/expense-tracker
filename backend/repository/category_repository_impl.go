package repository

import (
	"database/sql"

	"github.com/dnwandana/expense-tracker/model/entity"
	"github.com/dnwandana/expense-tracker/utils"
)

type CategoryRepositoryImpl struct {
	DB *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return &CategoryRepositoryImpl{DB: db}
}

func (repo *CategoryRepositoryImpl) Create(category *entity.Category) {
	id := utils.GenerateNanoID(5)
	_, err := repo.DB.Exec("INSERT INTO categories (id, user_id, name) VALUES (?, ?, ?)", id, category.UserID, category.Name)
	utils.PanicIfError(err)
}

func (repo *CategoryRepositoryImpl) FindByID(id string) *entity.Category {
	query := "SELECT id, name FROM categories WHERE id = ? ORDER BY name ASC"
	row, err := repo.DB.Query(query, id)
	utils.PanicIfError(err)
	defer row.Close()

	category := new(entity.Category)
	if row.Next() {
		err = row.Scan(&category.ID, &category.Name)
		utils.PanicIfError(err)
	}

	return category
}

func (repo *CategoryRepositoryImpl) FindByUserID(categoryID string) []*entity.Category {
	query := "SELECT id, name, created_at, updated_at FROM categories WHERE user_id = ?"
	rows, err := repo.DB.Query(query, categoryID)
	utils.PanicIfError(err)
	defer rows.Close()

	categories := make([]*entity.Category, 0)
	for rows.Next() {
		category := new(entity.Category)
		err = rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		utils.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}

func (repo *CategoryRepositoryImpl) Update(userID string, category *entity.Category) {
	_, err := repo.DB.Exec("UPDATE categories SET name = ? WHERE user_id = ? AND id = ?", category.Name, userID, category.ID)
	utils.PanicIfError(err)
}

func (repo *CategoryRepositoryImpl) Delete(userID, categoryID string) {
	_, err := repo.DB.Exec("DELETE FROM categories WHERE user_id = ? AND id = ?", userID, categoryID)
	utils.PanicIfError(err)
}