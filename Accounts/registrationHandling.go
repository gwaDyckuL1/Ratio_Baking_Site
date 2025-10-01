package accounts

import "database/sql"

func CheckEmail(email string, db *sql.DB) (bool, error) {
	var exists bool

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE email = ?
		)
	`
	err := db.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func CheckUserName(username string, db *sql.DB) (bool, error) {
	var exists bool

	query := `
		SELECT EXISTS (
			SELECT 1
			FROM users
			WHERE username = ?
		)
	`
	err := db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
