package handlers

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	accounts "github.com/gwaDyckuL1/Ratio_Baking_Site/Accounts"
)

func SessionMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionInfo := accounts.ActiveSession(db, r)

		if sessionInfo.LoggedIn {
			query := `UPDATE sessions 
				SET last_active = CURRENT_TIMESTAMP	
				WHERE session_token = ?;`

			_, err := db.Exec(query, sessionInfo.SessionToken)
			if err != nil {
				log.Println("Error in updating last_active: ", err)
			}
		}

		ctx := context.WithValue(r.Context(), "session", sessionInfo)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
