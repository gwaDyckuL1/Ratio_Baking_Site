package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func LogoutHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionInfo := r.Context().Value("session").(*models.Session)

		if !sessionInfo.LoggedIn {
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "session-token",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		})

		query := `
			DELETE FROM sessions
			WHERE session_token = ?;
		`
		_, err := db.Exec(query, sessionInfo.SessionToken)
		if err != nil {
			log.Printf("Error in removing session token %v from session: %v", sessionInfo.SessionToken, err)
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
