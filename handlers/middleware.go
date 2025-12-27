package handlers

import (
	"context"
	"database/sql"
	"net/http"

	accounts "github.com/gwaDyckuL1/Ratio_Baking_Site/Accounts"
)

func SessionMiddleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionInfo := accounts.ActiveSession(db, r)
		ctx := context.WithValue(r.Context(), "session", sessionInfo)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
