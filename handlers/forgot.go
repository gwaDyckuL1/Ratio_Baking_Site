package handlers

import (
	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
	"html/template"
	"net/http"
)

func ForgotLoginHandler(templates map[string]*template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionInfo := r.Context().Value("session").(*models.Session)
		err := templates["forgotLogin"].Execute(w, models.WebData{Session: sessionInfo})
		if err != nil {
			http.Error(w, "Template Error", http.StatusInternalServerError)
		}
	}
}
