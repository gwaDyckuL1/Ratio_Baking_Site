package accounts

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func ActiveSession(db *sql.DB, r *http.Request) models.Session {
	var userID string
	var s models.Session

	sessionToken, err := r.Cookie("session-token")
	if err != nil {
		s.LoggedIn = false
	}

	err = db.QueryRow(`
		SELECT user_id 
		FROM sessions 
		WHERE session_token = ?`, sessionToken.Value).Scan(&userID)
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
		WHERE id = ?
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
