package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectarBD() *sql.DB {
	conexao := "user=postgres dbname=oxidb password=yourpassword host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic("Nao conseguiu conectar com o banco de dados." + err.Error())
	}
	return db
}
