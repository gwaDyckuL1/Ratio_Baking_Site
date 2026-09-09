package calculator

import (
	"strconv"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func Calculator(data *models.RecipeData, problems models.FormErrors) {
	if data.Calculator == "bread" {
		breadCalculator(data, problems)
	}
	if data.Calculator == "cake" {
		cakeCalculator(data, problems)
	}
}

func cakeCalculator(data *models.RecipeData, problems models.FormErrors) {
	var eggPercent, fatPercent, hydrationPercent, saltPercent, sugarPercent float64

	if data.SubCalculator == "free-form" {
		eggPercent = stringToFloat("egg", data.EggIn, problems)
		fatPercent = stringToFloat("fat", data.FatIn, problems)
		hydrationPercent = stringToFloat("hydration", data.HydrationIn, problems)
		saltPercent = stringToFloat("salt", data.SaltIn, problems)
		sugarPercent = stringToFloat("sugar", data.SugarIn, problems)
	}

	var area float64
	if data.Shape == "weird" {
		area = stringToFloat("volumn", data.Volumn, problems)
	} else {
		area = getArea(data, problems)
	}
	totalWeight := getTotalWeight(area, data, problems)

	switch data.SubCalculator {
	case "free-form":
		data.DoughWeight = strconv.Itoa(totalWeight)
		getFlourWeight(data, problems)
		flour := stringToFloat("flour", data.FlourIn, problems)
		data.FlourOut = strconv.FormatFloat(flour, 'f', 0, 64)
		data.EggGrams = strconv.FormatFloat(flour*eggPercent/100, 'f', 0, 64)
		data.EggWhole = strconv.FormatFloat(flour*eggPercent/100/56, 'f', 1, 64) //56 is the average weight of a large egg
		data.FatOut = strconv.FormatFloat(flour*fatPercent/100, 'f', 0, 64)
		data.HydrationOut = strconv.FormatFloat(flour*hydrationPercent/100, 'f', 0, 64)
		data.SaltOut = strconv.FormatFloat(flour*saltPercent/100, 'f', 0, 64)
		data.SugarOut = strconv.FormatFloat(flour*sugarPercent/100, 'f', 0, 64)
	case "high-ratio-cake":
		fat := .2 * float64(totalWeight)
		eggGrams := fat
		subWeight := float64(totalWeight) - fat
		flour := subWeight / 3
		liquid := flour - eggGrams
		data.FlourOut = strconv.FormatFloat(flour, 'f', 0, 64)
		data.FatOut = strconv.FormatFloat(fat, 'f', 0, 64)
		data.SugarOut = strconv.FormatFloat(flour, 'f', 0, 64)
		data.EggGrams = strconv.FormatFloat(eggGrams, 'f', 0, 64)
		data.EggWhole = strconv.FormatFloat(eggGrams/56, 'f', 0, 64)
		data.HydrationOut = strconv.FormatFloat(liquid, 'f', 0, 64)
		data.SaltOut = strconv.FormatFloat(flour*0.02, 'f', 0, 64)
	case "pound-cake": //ratio is 1:1:1:1, flour, sugar, egg, fat
		flour := float64(totalWeight) / 4
		data.FlourOut = strconv.FormatFloat(flour, 'f', 0, 64)
		data.FatOut = strconv.FormatFloat(flour, 'f', 0, 64)
		data.SugarOut = strconv.FormatFloat(flour, 'f', 0, 64)
		data.EggGrams = strconv.FormatFloat(flour, 'f', 0, 64)
		data.EggWhole = strconv.FormatFloat(flour/56, 'f', 0, 64)
		data.SaltOut = strconv.FormatFloat(flour*0.02, 'f', 0, 64)
	case "sponge-cake": //ratio is 2:1:1, egg, flour, sugar
		flour := float64(totalWeight) / 4
		data.FlourOut = strconv.FormatFloat(flour, 'f', 0, 64)
		data.SugarOut = strconv.FormatFloat(flour, 'f', 0, 64)
		data.EggGrams = strconv.FormatFloat(flour*2, 'f', 0, 64)
		data.EggWhole = strconv.FormatFloat((flour*2)/56, 'f', 0, 64)
		data.SaltOut = strconv.FormatFloat(flour*0.02, 'f', 0, 64)
	}
}
