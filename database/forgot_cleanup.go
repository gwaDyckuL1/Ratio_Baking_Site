package database

import (
	"database/sql"
	"log"
)

func ForgotPasswordCleanup(db *sql.DB) {
	query := `
		DELETE
			FROM forgotToken
			WHERE created_at < DATETIME('NOW', '-30 MINUTES')
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Println("Error in deleting old tokens based on time", err)
	}
}
