package calculator

import (
	"strconv"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func breadCalculator(data *models.RecipeData, problems models.FormErrors) {
	if data.SubCalculator == "pan-dimension" {
		area := getArea(data, problems)
		totalWeight := getTotalWeight(area, data, problems)
		data.DoughWeight = strconv.Itoa(totalWeight)
	}
	if len(data.DoughWeight) > 0 {
		getFlourWeight(data, problems)
	}
	flour := stringToFloat("flour", data.FlourIn, problems)
	eggPercent := stringToFloat("egg", data.EggIn, problems)
	fatPercent := stringToFloat("fat", data.FatIn, problems)
	hydrationPercent := stringToFloat("hydration", data.HydrationIn, problems)
	saltPercent := stringToFloat("salt", data.SaltIn, problems)
	sugarPercent := stringToFloat("sugar", data.SugarIn, problems)
	leavenPercentage := 0.00
	if len(data.SourdoughIn) > 0 {
		leavenPercentage = stringToFloat("Sourdough", data.SourdoughIn, problems)
	} else {
		leavenPercentage = stringToFloat("Yeast", data.YeastIn, problems)
	}

	data.FlourOut = strconv.FormatFloat(flour, 'f', 0, 64)
	data.EggGrams = strconv.FormatFloat(flour*eggPercent/100, 'f', 0, 64)
	data.EggWhole = strconv.FormatFloat(flour*eggPercent/100/56, 'f', 1, 64) //56 is the average weight of a large egg
	data.FatOut = strconv.FormatFloat(flour*fatPercent/100, 'f', 0, 64)
	data.HydrationOut = strconv.FormatFloat(flour*hydrationPercent/100, 'f', 0, 64)
	data.SaltOut = strconv.FormatFloat(flour*saltPercent/100, 'f', 0, 64)
	data.SugarOut = strconv.FormatFloat(flour*sugarPercent/100, 'f', 0, 64)
	data.LeavenerOut = strconv.FormatFloat(flour*leavenPercentage/100, 'f', 0, 64)

	if len(data.TangzhongPercentage) > 0 {
		modifyForTangzhong(data, problems)
	}
}

// func calculateIngredients(data *models.RecipeData, problems models.FormErrors) {
// 	eggPercent := stringToFloat("egg", data.EggIn, problems)
// 	fatPercent := stringToFloat("fat", data.FatIn, problems)
// 	hydrationPercent := stringToFloat("hydration", data.HydrationIn, problems)
// 	saltPercent := stringToFloat("salt", data.SaltIn, problems)
// 	sugarPercent := stringToFloat("sugar", data.SugarIn, problems)

// 	if len(data.SaltIn) == 0 {
// 		saltPercent = 2
// 	}

// 	leaveningAmount := 0.00
// 	switch data.Leavener {
// 	case "Sourdough":
// 		if data.SourdoughIn == "" {
// 			leaveningAmount = 20
// 		} else {
// 			leaveningAmount = stringToFloat("sourdough", data.SourdoughIn, problems)
// 		}
// 	case "Yeast":
// 		if data.YeastIn == "" {
// 			leaveningAmount = 1
// 		} else {
// 			leaveningAmount = stringToFloat("yeast", data.YeastIn, problems)
// 		}
// 	}

// 	flour := 0.00
// 	if len(data.DoughWeight) > 0 {
// 		doughWeight := stringToFloat("Total Dough Weight", data.DoughWeight, problems)
// 		totalPercent := 100 + hydrationPercent + eggPercent + fatPercent + sugarPercent + saltPercent + leaveningAmount
// 		flour = (100 / totalPercent) * doughWeight
// 	} else {
// 		flour = stringToFloat("flour", data.FlourIn, problems)
// 	}

// 	tFlour, tHydration := 0.00, 0.00
// 	tangzhongPercentage := stringToFloat("tangzhong percentage", data.TangzhongPercentage, problems)
// 	if tangzhongPercentage > 0 {
// 		tRatio := stringToFloat("tangzhong ratio", data.TanghzhongRatio, problems)
// 		tFlour = (tangzhongPercentage / 100) / (1 + tRatio) * flour
// 		tHydration = tFlour * tRatio
// 		data.TangzhongFlour = fmt.Sprintf("%.0f", tFlour)
// 		data.TangzhongHydration = fmt.Sprintf("%.0f", tHydration)
// 	}

// 	eggGrams := (eggPercent / 100) * flour
// 	data.EggGrams = fmt.Sprintf("%.0f", eggGrams)
// 	data.EggWhole = fmt.Sprintf("%.1f", eggGrams/56.00)
// 	data.FatOut = fmt.Sprintf("%.0f", flour*fatPercent/100)
// 	data.FlourOut = fmt.Sprintf("%.0f", flour-tFlour)
// 	data.HydrationOut = fmt.Sprintf("%.0f", flour*hydrationPercent/100-tHydration)
// 	data.SaltOut = fmt.Sprintf("%.0f", flour*(saltPercent/100))
// 	data.SugarOut = fmt.Sprintf("%.0f", flour*(sugarPercent/100))
// 	data.LeavenerOut = fmt.Sprintf("%.0f", flour*(leaveningAmount/100))
// }
