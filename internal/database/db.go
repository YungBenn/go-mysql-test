package database

import (
	"database/sql"
	"os"
)

func ConnectDB() *sql.DB {
	dbUrl := os.Getenv("DB_URL")

	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}

	return db
}
