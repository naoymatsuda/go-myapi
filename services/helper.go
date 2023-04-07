package services

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./sampleDB.sql")
	if err != nil {
		return nil, err
	}
	return db, nil
}
