package handlers

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

func PasswordResetHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.URL.Query().Get("token")

		data := map[string]string{
			"Token": token,
		}

		tmpl := template.Must(template.ParseFiles(
			"templates/layout.html",
			"templates/passwordReset",
		))

		err := tmpl.Execute(w, data)
		if err != nil {
			log.Println(w, "Template error", http.StatusInternalServerError)
		}

	}
}
