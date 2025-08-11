package calculator

import (
	"fmt"
	"strconv"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func Calculator(data *models.RecipeData, problems models.FormErrors) {
	if data.Calculator == "bread" {
		breadCalculator(data, problems)
	}
}

func breadCalculator(data *models.RecipeData, problems models.FormErrors) {

	fat := stringToInt("fat", data.Fat, problems)
	hydration := stringToInt("hydration", data.Hydration, problems)
	salt := stringToInt("salt", data.Salt, problems)
	sugar := stringToInt("sugar", data.Sugar, problems)

	tangzhongPercentage := 0
	if len(data.TangzhongPercentage) > 0 {
		tangzhongPercentage = stringToInt("Tanghzhong Percentage", data.TangzhongPercentage, problems)
	}

	if data.SubCalculator == "flour-weight" {
		flour := stringToInt("flour", data.Flour, problems)
	}
	if data.SubCalculator == "total-weight" {
		doughWeight := stringToInt("Total Dough Weight", data.DoughWeight, problems)

	}
	if data.SubCalculator == "pan-dimension" {

	}
}

func stringToInt(name, strNum string, problems models.FormErrors) int {
	if len(strNum) == 0 {
		return 0
	}

	newInt, err := strconv.Atoi((strNum))
	if err != nil {
		problems[name] = fmt.Sprintf("The number %s for the field %s failed to convert", strNum, name)
		return 0
	}
	return newInt
}
