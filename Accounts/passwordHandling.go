package accounts

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(input, saved string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(input), []byte(saved))
	return err == nil
}

func GetPassword(username string, db *sql.DB) (string, error) {
	var hash string
	err := db.QueryRow("select password from users where username = ?", username).Scan(&hash)
	return hash, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
