package calculator

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func Calculator(data *models.RecipeData, problems models.FormErrors) {
	if data.Calculator == "bread" {
		breadCalculator(data, problems)
	}
}

func breadCalculator(data *models.RecipeData, problems models.FormErrors) {

	egg := stringToFloat("egg", data.EggIn, problems)
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
		eggGrams := (egg / 100) * flour
		data.EggGrams = fmt.Sprintf("%.0f", eggGrams)
		data.EggWhole = fmt.Sprintf("%.1f", eggGrams/56.00)
		data.FatOut = fmt.Sprint(int((fat / 100) * flour))
		data.HydrationOut = fmt.Sprint(int((hydration / 100) * flour))
		data.SaltOut = fmt.Sprint(int((salt / 100) * flour))
		data.SugarOut = fmt.Sprint(int((sugar / 100) * flour))
		data.LeavenerOut = fmt.Sprint(int((float64(leaveningAmount) / 100) * float64(flour)))

		tangzhongCheck(data, tangzhongPercentage, problems)
	}
	if data.SubCalculator == "total-weight" {
		doughWeight := stringToFloat("Total Dough Weight", data.DoughWeight, problems)
		totalPercent := 100 + hydration + fat + sugar + salt + leaveningAmount
		flour := (100 / totalPercent) * doughWeight
		eggGrams := (egg / 100) * flour
		data.EggGrams = fmt.Sprintf("%.0f", eggGrams)
		data.EggWhole = fmt.Sprintf("%.1f", eggGrams/56.00)
		data.FatOut = fmt.Sprintf("%.0f", flour*fat/100)
		data.FlourIn = fmt.Sprintf("%.0f", flour)
		data.FlourOut = fmt.Sprintf("%.0f", flour)
		data.HydrationOut = fmt.Sprintf("%.0f", flour*hydration/100)
		data.LeavenerOut = fmt.Sprintf("%.0f", flour*(leaveningAmount/100))
		data.SaltOut = fmt.Sprintf("%.0f", flour*(salt/100))
		data.SugarOut = fmt.Sprintf("%.0f", flour*(sugar/100))

		tangzhongCheck(data, tangzhongPercentage, problems)
	}
	if data.SubCalculator == "pan-dimension" {
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
		volumn := 0.00
		if data.Shape == "square" {
			volumn = height * width * depth
		} else {
			radius := diameter / 2
			volumn = math.Pi * math.Pow(radius, 2) * height
		}
		fmt.Println("Volumn is: ", volumn)
		//total dough weight = volumn * density * fill
		//starting with density of 1, like water
		// fill is how full the pan should be before proofing and baking
		doughWeight := 0.00
		if fat > 0 {
			doughWeight = volumn * 1.00 * 0.6
		} else {
			doughWeight = volumn * 1.00 * 0.5
		}
		fmt.Println("Total dough weight: ", doughWeight)
		totalPercent := 100 + hydration + fat + sugar + salt + leaveningAmount
		flour := (100 / totalPercent) * doughWeight
		eggGrams := (egg / 100) * flour
		data.EggGrams = fmt.Sprintf("%.0f", eggGrams)
		data.EggWhole = fmt.Sprintf("%.1f", eggGrams/56.00)
		data.FatOut = fmt.Sprintf("%.0f", flour*fat/100)
		data.FlourIn = fmt.Sprintf("%.0f", flour)
		data.FlourOut = fmt.Sprintf("%.0f", flour)
		data.HydrationOut = fmt.Sprintf("%.0f", flour*hydration/100)
		data.SaltOut = fmt.Sprintf("%.0f", flour*(salt/100))
		data.SugarOut = fmt.Sprintf("%.0f", flour*(sugar/100))
		data.LeavenerOut = fmt.Sprintf("%.0f", flour*(leaveningAmount/100))

		tangzhongCheck(data, tangzhongPercentage, problems)
	}
}

func tangzhongCheck(data *models.RecipeData, tangzhongPercentage float64, problems models.FormErrors) {
	if tangzhongPercentage > 0 {
		flour := stringToFloat("flour", data.FlourIn, problems)
		hydrationPercent := stringToFloat("hydration", data.HydrationIn, problems)
		hydration := flour * hydrationPercent / 100
		tRatio := stringToFloat("tangzhong ratio", data.TanghzhongRatio, problems)

		tFlour := (tangzhongPercentage / 100) / (1 + tRatio) * flour
		tHydration := tFlour * tRatio

		data.TangzhongFlour = fmt.Sprintf("%.0f", tFlour)
		data.TangzhongHydration = fmt.Sprintf("%.0f", tHydration)
		data.FlourOut = fmt.Sprintf("%.0f", flour-tFlour)
		data.HydrationOut = fmt.Sprintf("%.0f", hydration-tHydration)
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
