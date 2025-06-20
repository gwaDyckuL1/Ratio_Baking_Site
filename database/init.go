package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "database/ratio.db")
	if err != nil {
		log.Fatal("Failed to open database. Error:", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		username TEXT NOT NULL UNIQUE,
		name TEXT,
		email TEXT NOT NULL,
		password TEXT NOT NULL
	);
	`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Error creating user table:", err)
	}
	return db
}
