package repository

import "github.com/dnwandana/expense-tracker/model/entity"

type CategoryRepository interface {
	Create(category *entity.Category)
	FindByID(categoryID int) *entity.Category
	FindByUserID(userID int) []*entity.Category
	Update(userID int, category *entity.Category)
	Delete(userID, categoryID int)
}
