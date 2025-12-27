package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	accounts "github.com/gwaDyckuL1/Ratio_Baking_Site/Accounts"
	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func RegistrationSubmitHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		sessionInfo := r.Context().Value("session").(*models.Session)
		if r.Method != http.MethodPost || sessionInfo.LoggedIn {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		err := r.ParseMultipartForm(4 << 20)
		if err != nil {
			http.Error(w, "Invalid form submission", http.StatusBadRequest)
			return
		}

		data := models.RegistrationData{
			Username: r.FormValue("username"),
			Name:     r.FormValue("name"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		emailUsed, err := accounts.CheckEmail(data.Email, db)
		if err != nil && err != sql.ErrNoRows {
			log.Printf("Database error checking email: %v", err)
			http.Error(w, "Internal Server Error. Please try again later", http.StatusInternalServerError)
			return
		}
		if emailUsed {
			json.NewEncoder(w).Encode(models.Response{
				Ok:      false,
				Field:   "email",
				Message: "This email already has an account.",
			})
			return
		}

		usernameUsed, err := accounts.CheckUserName(data.Username, db)
		if err != nil && err != sql.ErrNoRows {
			log.Printf("Database error checking username: %v", err)
			http.Error(w, "Internal Server Error. Please try again later", http.StatusInternalServerError)
			return
		}
		if usernameUsed {
			json.NewEncoder(w).Encode(models.Response{
				Ok:      false,
				Field:   "username",
				Message: "Username not available. Please choose another.",
			})
			return
		}

		hashPassword, err := accounts.HashPassword(data.Password)
		if err != nil {
			log.Printf("Error in hashing password: %v", err)
			http.Error(w, "Internal Server Error. Please try again later", http.StatusInternalServerError)
			return
		}

		_, err = db.Exec(`INSERT INTO 
			users (username, name, email, password, role, create_date)
			VALUES (?, ?, ?, ?, ?, DATETIME("NOW"));`,
			data.Username, data.Name, data.Email, hashPassword, "User")

		if err != nil {
			log.Printf("Error inserting new user in database. Error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(models.Response{Ok: false, Message: "Server error. Try again later."})
			return
		}

		json.NewEncoder(w).Encode(models.Response{Ok: true, Message: "Registration Successful"})

	}
}
