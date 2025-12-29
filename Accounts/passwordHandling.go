package accounts

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(input, saved string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(saved), []byte(input))
	return err == nil
}

func GetPassword(username string, db *sql.DB) (int, string, error) {
	var hash string
	var id int
	err := db.QueryRow("select id, password from users where username = ?;", username).Scan(&id, &hash)
	return id, hash, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
