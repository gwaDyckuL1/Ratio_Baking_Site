package calculator

import (
	"fmt"
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
