package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaPostgres() *sql.DB {
	conexao := "user=postgres dbname=go password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
