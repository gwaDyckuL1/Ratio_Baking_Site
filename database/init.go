package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	dbPath := os.Getenv("DATABASE_PATH")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open database. Error:", err)
	}

	createTableSQL := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			name TEXT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			create_date DATETIME DEFAULT CURRENT_TIMESTAMP,
			last_login DATETIME,
			role TEXT
		);`,
		`CREATE TABLE IF NOT EXISTS recipe (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			recipe_name TEXT NOT NULL,
			recipe_data TEXT NOT NULL,
			create_date DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)`,
	}

	for _, table := range createTableSQL {
		_, err = db.Exec(table)
		if err != nil {
			log.Fatal("Error creating user table:", err)
		}
	}
	return db
}
