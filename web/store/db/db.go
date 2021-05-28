package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectDatabase() *sql.DB {
	connection := "postgres://postgres:postgres@localhost/go_store?sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	return db
}


