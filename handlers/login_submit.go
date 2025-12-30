package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	accounts "github.com/gwaDyckuL1/Ratio_Baking_Site/Accounts"
	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func LoginSubmitHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		sessionInfo := r.Context().Value("session").(*models.Session)
		if sessionInfo.LoggedIn {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		err := r.ParseMultipartForm(4 << 20)
		if err != nil {
			http.Error(w, "Error parsing form.", http.StatusBadRequest)
			return
		}
		data := models.Login{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		userId, savedPassword, err := accounts.GetPassword(data.Username, db)
		if err != nil {
			if err == sql.ErrNoRows {
				json.NewEncoder(w).Encode(models.Response{
					Ok:      false,
					Field:   "login-error",
					Message: "The username or password is incorrect.",
				})
			} else {
				json.NewEncoder(w).Encode(models.Response{
					Ok:      false,
					Field:   "login-error",
					Message: "Internal failure. Please try again later",
				})
				log.Println("Login DB error for username: ", data.Username, err)
			}
			return
		}

		passwordGood := accounts.CheckPassword(data.Password, savedPassword)
		if passwordGood {
			sessionID := accounts.NewSessionID()

			_, err = db.Exec(`
				INSERT INTO sessions (user_id, session_token)
				VALUES (?, ?);
				`, userId, sessionID)
			if err != nil {
				log.Printf("Error in saving session cookie. %v", err)
			}

			_, err = db.Exec(`
				UPDATE users 
				SET last_login = CURRENT_TIMESTAMP
				WHERE id = ?;
				`, userId)
			if err != nil {
				log.Printf("Error saving last login for %v. %v", data.Username, err)
			}

			http.SetCookie(w, &http.Cookie{
				Name:     "session-token",
				Value:    sessionID,
				Path:     "/",
				HttpOnly: true,
				Secure:   false,
				SameSite: http.SameSiteLaxMode,
			})

		} else {
			json.NewEncoder(w).Encode(models.Response{
				Ok:      false,
				Field:   "login-error",
				Message: "The username or password is incorrect.",
			})
			return
		}

		json.NewEncoder(w).Encode(models.Response{Ok: true})

	}
}
