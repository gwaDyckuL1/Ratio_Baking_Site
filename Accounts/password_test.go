package accounts

import (
	"testing"
)

func TestHashing(t *testing.T) {
	password := "1qaz@WSX"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Problem hashing password")
	}

	match := CheckPassword(password, hashedPassword)
	if !match {
		t.Errorf("CheckPassword not working")
	}
}
