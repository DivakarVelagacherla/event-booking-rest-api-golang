package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("DB Connection error")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
