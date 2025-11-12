package accounts

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func ActiveSession(db *sql.DB, sessionToken string) models.Session {
	var userID string
	var s models.Session

	err := db.QueryRow(`
		SELECT user_id 
		FROM sessions 
		WHERE session_token = ?`, sessionToken).Scan(&userID)
	if err == sql.ErrNoRows {
		fmt.Println("No session found.")
		return models.Session{
			LoggedIn: false,
		}
	}
	if err != nil {
		fmt.Println("Error in getting User Id: ", err)
		return models.Session{LoggedIn: false}
	}

	err = db.QueryRow(`
		SELECT name, username
		FROM users
		WHERE user_id = ?
	`, userID).Scan(&s.Name, &s.Username)
	if err == sql.ErrNoRows {
		fmt.Println("User Id not found")
		return models.Session{LoggedIn: false}
	}
	if err != nil {
		fmt.Println("Error in getting User info: ", err)
		return models.Session{LoggedIn: false}
	}

	s.LoggedIn = true
	return s
}

func NewSessionID() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}
