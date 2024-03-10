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
	id := utils.GenerateNanoID(5)
	_, err := repo.DB.Exec("INSERT INTO users (id, username, password) VALUES (?, ?, ?)", id, user.Username, user.Password)
	utils.PanicIfError(err)
}

func (repo *UserRepositoryImpl) FindByUsername(username string) *entity.User {
	user := new(entity.User)
	rows, err := repo.DB.Query("SELECT id, username, password FROM users WHERE username = ?", username)
	utils.PanicIfError(err)
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password)
		utils.PanicIfError(err)
	}

	return user
}
