package repository

import "github.com/dnwandana/expense-tracker/model/entity"

type CategoryRepository interface {
	Create(category *entity.Category)
	FindByID(categoryID string) *entity.Category
	FindByUserID(userID string) []*entity.Category
	Update(userID string, category *entity.Category)
	Delete(userID, categoryID string)
}
