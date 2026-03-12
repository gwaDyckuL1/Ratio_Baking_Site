package accounts

import (
	"crypto/sha256"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CheckPassword(input, saved string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(saved), []byte(input))
	return err == nil
}

func GetPassword(email string, db *sql.DB) (int, string, error) {
	var hash string
	var id int
	err := db.QueryRow("select id, password from users where email = ?;", email).Scan(&id, &hash)
	return id, hash, err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func HashToken(token string) string {
	h := sha256.New()
	h.Write([]byte(token))
	return fmt.Sprintf("%x", h.Sum(nil))
}
