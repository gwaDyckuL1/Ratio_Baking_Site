package accounts

import "database/sql"

func CheckUserName(username string, db *sql.DB) bool {
	var exists bool

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE username = ?
		)
	`
	err := db.QueryRow(query, username).Scan(&exists)

}
