package handlers

import (
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	accounts "github.com/gwaDyckuL1/Ratio_Baking_Site/Accounts"
	"github.com/gwaDyckuL1/Ratio_Baking_Site/database"
	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func ForgotLoginSubmitHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		database.ForgotPasswordCleanup(db)

		err := r.ParseMultipartForm(4 << 20)
		if err != nil {
			log.Printf("Error parsing forgot password form: %v", err)
			http.Error(w, "Error parsing form.", http.StatusBadRequest)
			return
		}
		email := r.FormValue("email")
		query := `
			SELECT username
			FROM users
			WHERE UPPER(email) = UPPER(?)
		  `
		var username string
		err = db.QueryRow(query, email).Scan(&username)
		if err != nil {
			json.NewEncoder(w).Encode(models.Response{Ok: true})
			return
		}

		token := rand.Text()
		hashToken := accounts.HashToken(token)
		if err != nil {
			log.Printf("Error in hashing token: %v", err)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(models.Response{Ok: true})
			return
		}
		query = `
			INSERT INTO forgotToken (token, email)
			VALUES (?,?)
		`
		_, err = db.Exec(query, hashToken, email)
		if err != nil {
			log.Printf("Error inserting token: %v", err)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(models.Response{Ok: true})
			return
		}

		resetURL := fmt.Sprintf("https://localhost:8080/reseetPassword?token=%s", token)
		fmt.Printf("To resert your password follow this link: %s", resetURL)

		json.NewEncoder(w).Encode(models.Response{Ok: true})
	}
}
