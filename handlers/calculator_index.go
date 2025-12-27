package handlers

import (
	"html/template"
	"net/http"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func CalculatorIndexHandler(templates map[string]*template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionInfo := r.Context().Value("session").(*models.Session)
		err := templates["calcIndex"].Execute(w, models.WebData{Session: sessionInfo})
		if err != nil {
			http.Error(w, "Template Error", http.StatusInternalServerError)
		}
	}
}
