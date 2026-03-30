package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	accounts "github.com/gwaDyckuL1/Ratio_Baking_Site/Accounts"
	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func PasswordResetSubmit(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		err := r.ParseMultipartForm(4 << 20)
		if err != nil {
			json.NewEncoder(w).Encode(models.Response{
				Ok:      false,
				Field:   "token",
				Message: "Internal error, please try again leter.",
			})
			return
		}

		token := r.FormValue("token")
		hashToken := accounts.HashToken(token)
		var email string
		query := `
			SELECT email
			FROM forgotToken
			WHERE token = ?;
		`
		err = db.QueryRow(query, hashToken).Scan(&email)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				json.NewEncoder(w).Encode(models.Response{
					Ok:      false,
					Field:   "token",
					Message: "Token has expired. Please resubmit request",
				})
				return
			}
			log.Printf("Error using token to get email: %v", err)
			json.NewEncoder(w).Encode(models.Response{
				Ok:      false,
				Field:   "token",
				Message: "Internal error, please try again leter.",
			})
			return
		}

		password := r.FormValue("password")
		hashPassword, err := accounts.HashPassword(password)
		if err != nil {
			log.Printf("Error in hashing password: %v", err)
			json.NewEncoder(w).Encode(models.Response{
				Ok:      false,
				Field:   "token",
				Message: "Internal error, please try again leter.",
			})
			return
		}

		query = `
			UPDATE users
			SET password = ?
			WHERE email = ?;
		`
		_, err = db.Exec(query, hashPassword, email)
		if err != nil {
			log.Printf("Error in updating password for %v: %v", email, err)
			json.NewEncoder(w).Encode(models.Response{
				Ok:      false,
				Field:   "token",
				Message: "Internal error, please try again leter.",
			})
			return
		}

		query = `
			DELETE
			FROM forgotToken
			WHERE token = ?;
		`
		_, err = db.Exec(query, hashToken)
		if err != nil {
			log.Printf("Error removing token after changing password: %v", err)
		}

		json.NewEncoder(w).Encode(models.Response{
			Ok:      true,
			Field:   "message",
			Message: "Password successfully changed.",
		})
	}
}
