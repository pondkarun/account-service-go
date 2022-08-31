package Services

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Condb() *sql.DB {
	connStr := "postgres://postgres:root@localhost/accountdb"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
