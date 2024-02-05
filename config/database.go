package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Postgres goalng driver
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "secret"
	dbname   = "postgres"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		//panic(err)
	}

	return db
}
