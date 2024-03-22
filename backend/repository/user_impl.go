package repository

import (
	"database/sql"

	"github.com/dnwandana/expense-tracker/model/entity"
	"github.com/dnwandana/expense-tracker/utils"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{DB: db}
}

func (repo *UserRepositoryImpl) Create(user *entity.User) {
	_, err := repo.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	utils.PanicIfError(err)
}

func (repo *UserRepositoryImpl) FindByUsername(username string) *entity.User {
	rows, err := repo.DB.Query("SELECT id, username, password FROM users WHERE username = ?", username)
	utils.PanicIfError(err)
	defer rows.Close()

	user := new(entity.User)
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password)
		utils.PanicIfError(err)
	}

	return user
}
