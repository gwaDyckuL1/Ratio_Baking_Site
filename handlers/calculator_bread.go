package handlers

import (
	"html/template"
	"net/http"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func CalculatorBreadHandler(templates map[string]*template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionInfo := r.Context().Value("session").(*models.Session)
		err := templates["bread"].Execute(w, models.WebData{Session: sessionInfo})
		if err != nil {
			http.Error(w, "Template error", http.StatusInternalServerError)
		}
	}
}
