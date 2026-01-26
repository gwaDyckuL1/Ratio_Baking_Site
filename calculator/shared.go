package calculator

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

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

func getArea(data *models.RecipeData, problems models.FormErrors) float64 {
	height := stringToFloat("Pan Height", data.Height, problems)
	width := stringToFloat("Pan Width", data.Width, problems)
	depth := stringToFloat("Pan Depth", data.Depth, problems)
	diameter := stringToFloat("Pan Radius", data.Diameter, problems)

	if data.Measurement == "inches" {
		//inch to centimeter conversion is 1 : 2.54
		height = height * 2.54
		width = width * 2.54
		depth = depth * 2.54
		diameter = diameter * 2.54
	}
	area := 0.00
	switch data.Shape {
	case "square":
		area = height * width * depth
	case "circle":
		radius := diameter / 2
		area = math.Pi * radius * radius * depth
	}
	return area
}

func getFlourWeight(data *models.RecipeData, problems models.FormErrors) {
	doughWeight := stringToFloat("Total Dough Weight", data.DoughWeight, problems)
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

	totalPercentage := 100 + eggPercent + fatPercent + hydrationPercent + saltPercent + sugarPercent + leavenPercentage
	flour := (100 / totalPercentage) * doughWeight
	data.FlourIn = strconv.FormatFloat(flour, 'f', -1, 64)
}

func getTotalWeight(area float64, data *models.RecipeData, problems models.FormErrors) int {
	fat := stringToFloat("fat", data.FatIn, problems)
	totalWeight := 0.00
	if fat > 0 {
		totalWeight = area * 1.00 * 0.6
	} else {
		totalWeight = area * 1.00 * 0.5
	}
	return int(math.Ceil(totalWeight))
}

func modifyForTangzhong(data *models.RecipeData, problems models.FormErrors) {
	flour := stringToFloat("flour", data.FlourOut, problems)
	hydration := stringToFloat("hydration", data.HydrationIn, problems)
	tangzhongPercentage := stringToFloat("tangzhong Percentage", data.TangzhongPercentage, problems)
	ratio := stringToFloat("Ratio", data.TanghzhongRatio, problems)

	tFlour := (tangzhongPercentage / 100) / (1 + ratio) * flour
	tHydration := tFlour * ratio

	data.TangzhongFlour = strconv.FormatFloat(tFlour, 'f', 0, 64)
	data.TangzhongHydration = strconv.FormatFloat(tHydration, 'f', 0, 64)
	data.FlourOut = strconv.FormatFloat(flour-tFlour, 'f', 0, 64)
	data.HydrationOut = strconv.FormatFloat(hydration-tHydration, 'f', 0, 64)
}
