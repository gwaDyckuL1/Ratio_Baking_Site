package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func DashboardDataHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Context-Type", "application/json")

		sessionInfo := r.Context().Value("session").(*models.Session)

		if !sessionInfo.LoggedIn {
			json.NewEncoder(w).Encode(models.Response{
				Ok: false,
			})
			return
		}

		type recipeToJS struct {
			Recipe_name string `json:"recipe_name"`
			Category    string `json:"category"`
			Subcategory string `json:"subcategory"`
			Recipe_data string `json:"recipe_data"`
			Notes       string `json:"notes"`
		}

		query := `SELECT
					recipe_name,
					category,
					subcategory,
					recipe_data,
					notes
				FROM recipe
				WHERE user_id = ?;
		`
		rows, err := db.Query(query, sessionInfo.UserID)
		if err != nil {
			log.Printf("Issue getting recipe data for %s. %v", sessionInfo.Name, err)
		}
		defer rows.Close()

		var recipeArray []recipeToJS
		for rows.Next() {

			var temp recipeToJS

			err := rows.Scan(
				&temp.Recipe_name,
				&temp.Category,
				&temp.Subcategory,
				&temp.Recipe_data,
				&temp.Notes,
			)
			if err != nil {
				log.Printf("Failed to scan row into struct: %v", err)
			}

			recipeArray = append(recipeArray, temp)
		}

		json.NewEncoder(w).Encode(recipeArray)
	}
}
