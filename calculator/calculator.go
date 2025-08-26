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

	fat := stringToFloat("fat", data.FatIn, problems)
	hydration := stringToFloat("hydration", data.HydrationIn, problems)
	salt := stringToFloat("salt", data.SaltIn, problems)
	sugar := stringToFloat("sugar", data.SugarIn, problems)

	if len(data.SaltIn) == 0 {
		salt = 2
	}

	leaveningAmount := 0.00
	switch data.Leavener {
	case "Sourdough":
		if data.SourdoughIn == "" {
			leaveningAmount = 20
		} else {
			leaveningAmount = stringToFloat("sourdough", data.SourdoughIn, problems)
		}
	case "Yeast":
		if data.YeastIn == "" {
			leaveningAmount = 1
		} else {
			leaveningAmount = stringToFloat("yeast", data.YeastIn, problems)
		}
	}

	tangzhongPercentage := 0.00
	if len(data.TangzhongPercentage) > 0 {
		tangzhongPercentage = stringToFloat("Tanghzhong Percentage", data.TangzhongPercentage, problems)
	}

	if data.SubCalculator == "flour-weight" {
		flour := stringToFloat("flour", data.FlourIn, problems)
		data.FlourOut = fmt.Sprint(flour)
		data.FatOut = fmt.Sprint(int((fat / 100) * flour))
		data.HydrationOut = fmt.Sprint(int((hydration / 100) * flour))
		data.SaltOut = fmt.Sprint(int((salt / 100) * flour))
		data.SugarOut = fmt.Sprint(int((sugar / 100) * flour))
		data.LeavenerOut = fmt.Sprint(int((float64(leaveningAmount) / 100) * float64(flour)))

		tangzhongCheck(data, tangzhongPercentage, problems)
	}
	if data.SubCalculator == "total-weight" {
		//doughWeight := stringToFloat("Total Dough Weight", data.DoughWeight, problems)

	}
	if data.SubCalculator == "pan-dimension" {

	}
}

func tangzhongCheck(data *models.RecipeData, tangzhongPercentage float64, problems models.FormErrors) {
	if tangzhongPercentage > 0 {
		flour := stringToFloat("flour", data.FlourIn, problems)
		hydration := stringToFloat("hydration", data.HydrationIn, problems)

		tFlour := (tangzhongPercentage / 100) * flour
		tHydration := tFlour * 5

		data.TangzhongFlour = fmt.Sprint(tFlour)
		data.TangzhongHydration = fmt.Sprint(tHydration)
		data.FlourOut = fmt.Sprint(flour - tFlour)
		data.HydrationOut = fmt.Sprint(hydration - tHydration)
	}
}

func stringToFloat(name, strNum string, problems models.FormErrors) float64 {
	if len(strNum) == 0 {
		return 0
	}

	newNum, err := strconv.ParseFloat(strNum, 64)
	if err != nil {
		problems[name] = fmt.Sprintf("The number %s for the field %s failed to convert", strNum, name)
		return 0
	}
	return newNum
}
