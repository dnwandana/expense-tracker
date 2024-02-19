package repository

import "github.com/dnwandana/expense-tracker/model/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindByUsername(username string) (*entity.User, error)
}
