package calculator

import (
	"fmt"
	"math"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func breadCalculator(data *models.RecipeData, problems models.FormErrors) {
	if data.SubCalculator == "pan-dimension" {
		panDimension(data, problems)
	} else {
		calculateIngredients(data, problems)
	}
}

func panDimension(data *models.RecipeData, problems models.FormErrors) {
	height := stringToFloat("Pan Height", data.Height, problems)
	width := stringToFloat("Pan Width", data.Width, problems)
	depth := stringToFloat("Pan Depth", data.Depth, problems)
	diameter := stringToFloat("Pan Radius", data.Diameter, problems)
	fat := stringToFloat("fat", data.FatIn, problems)
	if data.Measurement == "inches" {
		//inch to centimeter conversion is 1 : 2.54
		height = height * 2.54
		width = width * 2.54
		depth = depth * 2.54
		diameter = diameter * 2.54
	}
	volumn := 0.00
	if data.Shape == "square" {
		volumn = height * width * depth
	} else {
		radius := diameter / 2
		volumn = math.Pi * math.Pow(radius, 2) * depth
	}
	doughWeight := 0.00
	if fat > 0 {
		doughWeight = volumn * 1.00 * 0.6
	} else {
		doughWeight = volumn * 1.00 * 0.5
	}
	data.DoughWeight = fmt.Sprintf("%.0f", doughWeight)
	calculateIngredients(data, problems)
}

func calculateIngredients(data *models.RecipeData, problems models.FormErrors) {
	eggPercent := stringToFloat("egg", data.EggIn, problems)
	fatPercent := stringToFloat("fat", data.FatIn, problems)
	hydrationPercent := stringToFloat("hydration", data.HydrationIn, problems)
	saltPercent := stringToFloat("salt", data.SaltIn, problems)
	sugarPercent := stringToFloat("sugar", data.SugarIn, problems)

	if len(data.SaltIn) == 0 {
		saltPercent = 2
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

	flour := 0.00
	if len(data.DoughWeight) > 0 {
		doughWeight := stringToFloat("Total Dough Weight", data.DoughWeight, problems)
		totalPercent := 100 + hydrationPercent + eggPercent + fatPercent + sugarPercent + saltPercent + leaveningAmount
		flour = (100 / totalPercent) * doughWeight
	} else {
		flour = stringToFloat("flour", data.FlourIn, problems)
	}

	tFlour, tHydration := 0.00, 0.00
	tangzhongPercentage := stringToFloat("tangzhong percentage", data.TangzhongPercentage, problems)
	if tangzhongPercentage > 0 {
		tRatio := stringToFloat("tangzhong ratio", data.TanghzhongRatio, problems)
		tFlour = (tangzhongPercentage / 100) / (1 + tRatio) * flour
		tHydration = tFlour * tRatio
		data.TangzhongFlour = fmt.Sprintf("%.0f", tFlour)
		data.TangzhongHydration = fmt.Sprintf("%.0f", tHydration)
	}

	eggGrams := (eggPercent / 100) * flour
	data.EggGrams = fmt.Sprintf("%.0f", eggGrams)
	data.EggWhole = fmt.Sprintf("%.1f", eggGrams/56.00)
	data.FatOut = fmt.Sprintf("%.0f", flour*fatPercent/100)
	data.FlourOut = fmt.Sprintf("%.0f", flour-tFlour)
	data.HydrationOut = fmt.Sprintf("%.0f", flour*hydrationPercent/100-tHydration)
	data.SaltOut = fmt.Sprintf("%.0f", flour*(saltPercent/100))
	data.SugarOut = fmt.Sprintf("%.0f", flour*(sugarPercent/100))
	data.LeavenerOut = fmt.Sprintf("%.0f", flour*(leaveningAmount/100))
}
