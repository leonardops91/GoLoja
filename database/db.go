package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "postgres"
)

func DbConnect() *sql.DB {
	Psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", Psqlconn)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	return db
}
