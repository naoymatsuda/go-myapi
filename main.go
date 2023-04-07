package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/naoymatsuda/myapi/api"
)

func main() {
	db, err := sql.Open("sqlite3", "./sampleDB.sql")
	if err != nil {
		log.Println("fail to connect db")
		return
	}
	r := api.NewRouter(db)

	log.Println("server start at post 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
