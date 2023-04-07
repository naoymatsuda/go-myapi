package repositories_test

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var testDB *sql.DB

func setup() error {
	var err error
	testDB, err = sql.Open("sqlite3", "../sampleDB.sql")
	if err != nil {
		return err
	}
	return nil
}

func teardown() error {
	if err := testDB.Close(); err != nil {
		return err
	}
	return nil
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}
	m.Run()
	teardown()
}
