package accounts

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPassword(input, saved string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(input), []byte(saved))
	return err == nil
}
