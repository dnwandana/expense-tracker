package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/dnwandana/expense-tracker/utils"
	_ "github.com/go-sql-driver/mysql"
)

// NewConnection creates a new pool of connections to the database
func NewConnection() *sql.DB {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)

	db, err := sql.Open("mysql", dataSource)
	utils.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
