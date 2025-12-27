package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/calculator"
	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func CalcResultsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		sessionInfo := r.Context().Value("session").(*models.Session)

		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		data := models.RecipeData{
			Calculator:          r.FormValue("calculatorFor"),
			SubCalculator:       r.FormValue("calculator-bread"),
			Measurement:         r.FormValue("measurement"),
			Shape:               r.FormValue("shape"),
			Height:              r.FormValue("height"),
			Width:               r.FormValue("width"),
			Depth:               r.FormValue("depth"),
			Diameter:            r.FormValue("diameter"),
			FlourIn:             r.FormValue("flour"),
			DoughWeight:         r.FormValue("dough-weight"),
			HydrationIn:         r.FormValue("hydration"),
			EggIn:               r.FormValue("egg"),
			FatIn:               r.FormValue("fat"),
			SugarIn:             r.FormValue("sugar"),
			TangzhongPercentage: r.FormValue("tangzhong-percentage"),
			TanghzhongRatio:     r.FormValue("tangzhong-ratio"),
			SaltIn:              r.FormValue("salt"),
			Leavener:            r.FormValue("leavener-choice"),
			SourdoughIn:         r.FormValue("sourdough"),
			YeastIn:             r.FormValue("yeast"),
		}

		problems := models.FormErrors{}

		calculator.Calculator(&data, problems)

		webData := models.WebData{
			Session:    sessionInfo,
			RecipeData: &data,
		}

		tmpl := template.Must(template.ParseFiles(
			"templates/layout.html",
			"templates/calculator/layout.html",
			"templates/calculator/results.html",
		))

		err = tmpl.Execute(w, webData)
		if err != nil {
			log.Println(w, "Template error", http.StatusInternalServerError)
		}
	}
}
