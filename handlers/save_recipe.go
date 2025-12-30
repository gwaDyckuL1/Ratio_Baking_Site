package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func SaveRecipeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionInfo := r.Context().Value("session").(*models.Session)

		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		err := r.ParseMultipartForm(4 << 20)
		if err != nil {
			log.Println("Error parsing recipe to save. ", err)
			json.NewEncoder(w).Encode(models.Response{
				Ok:    false,
				Field: "save-error",
			})
			return
		}

		recipeName := r.FormValue("recipeName")
		recipeJSON := r.FormValue("recipeJSON")
		notes := r.FormValue("notes")

		var recipe models.RecipeData
		err = json.Unmarshal([]byte(recipeJSON), &recipe)
		if err != nil {
			log.Println("Problem unmarsheling recipe during save process.", err)
		}

		stmt := `INSERT INTO
		recipe(user_id, recipe_name, category, subcategory, recipe_data, notes)
		VALUES(?,?,?,?,?,?);`

		_, err = db.Exec(stmt, sessionInfo.UserID, recipeName, recipe.Calculator, recipe.SubCalculator, recipeJSON, notes)
		if err != nil {
			log.Println("Error in inserting recipe into database.", err)
			json.NewEncoder(w).Encode(models.Response{
				Ok:    false,
				Field: "save-error",
			})
		} else {
			json.NewEncoder(w).Encode(models.Response{
				Ok:    true,
				Field: "success",
			})
		}
	}
}
