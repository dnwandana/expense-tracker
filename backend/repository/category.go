package repository

import (
	"github.com/dnwandana/expense-tracker/model/entity"
	"github.com/dnwandana/expense-tracker/model/web/response"
)

type CategoryRepository interface {
	Create(category *entity.Category)
	FindOne(userID, categoryID int) *response.Category
	FindByUserID(userID int) []*response.Category
	Update(userID int, category *entity.Category)
	Delete(userID, categoryID int)
}
