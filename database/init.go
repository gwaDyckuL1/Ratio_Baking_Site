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

	createTableSQL := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			name TEXT,
			email TEXT NOT NULL,
			password TEXT NOT NULL,
			create_date DATETIME DEFAULT CURRENT_TIMESTAMP
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
