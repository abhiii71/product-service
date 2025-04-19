package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectPostgres(dbUrl string) *sql.DB {
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Postgres: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Cannot reach Postgres: %v", err)
	}
	return db
}
