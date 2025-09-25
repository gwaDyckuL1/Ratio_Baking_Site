package calculator

import (
	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func Calculator(data *models.RecipeData, problems models.FormErrors) {
	if data.Calculator == "bread" {
		breadCalculator(data, problems)
	}
}
