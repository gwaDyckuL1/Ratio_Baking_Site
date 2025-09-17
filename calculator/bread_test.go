package calculator

import (
	"testing"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
)

func TestFlourWeight(t *testing.T) {
	data := models.RecipeData{
		Calculator:    "bread",
		SubCalculator: "flour-weight",
		FlourIn:       "1000",
		EggIn:         "10",
		FatIn:         "10",
		SaltIn:        "10",
		SugarIn:       "10",
		HydrationIn:   "75",
	}

	problems := models.FormErrors{}

	Calculator(&data, problems)

	if data.FlourOut != "1000" {
		t.Errorf("FLOUR: Expecting: 1000, Got: %s", data.FlourOut)
	}
	if data.EggGrams != "100" {
		t.Errorf("EGG GRAMS: Expected: 100, Got: %s", data.EggGrams)
	}
	if data.EggWhole != "1.8" {
		t.Errorf("WHOLE EGGS: Expected: 1.8, Got: %s", data.EggWhole)
	}
	if data.FatOut != "100" {
		t.Errorf("FAT: Expected: 100, Got: %s", data.FatOut)
	}
	if data.SaltOut != "100" {
		t.Errorf("SALT: Expected: 100, Got: %s", data.SaltOut)
	}
	if data.SugarOut != "100" {
		t.Errorf("SUGAR: Expected: 100, Got: %s", data.SugarOut)
	}
	if data.HydrationOut != "750" {
		t.Errorf("HYDRATION: Expected 750, Got: %s", data.HydrationOut)
	}
}

func TestLeavener(t *testing.T) {
	tests := []struct {
		name           string
		leavener       string
		flour          string
		leavenerAmount string
		expected       string
	}{
		{"Sourdough Value", "Sourdough", "100", "10", "10"},
		{"Sourdough No Value", "Sourdough", "100", "", "20"},
		{"Yeast Value", "Yeast", "100", "10", "10"},
		{"Yeast No Value", "Yeast", "100", "", "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			data := models.RecipeData{
				Calculator:    "bread",
				SubCalculator: "flour-weight",
				Leavener:      tt.leavener,
				FlourIn:       tt.flour,
			}
			if data.Leavener == "Sourdough" {
				data.SourdoughIn = tt.leavenerAmount
			} else {
				data.YeastIn = tt.leavenerAmount
			}
			problems := models.FormErrors{}
			Calculator(&data, problems)

			if data.LeavenerOut != tt.expected {
				t.Errorf("Got: %s, Wanted: %s", data.LeavenerOut, tt.expected)
			}
		})
	}

}

func TestTangzhong(t *testing.T) {
	data := models.RecipeData{
		Calculator:          "bread",
		SubCalculator:       "flour-weight",
		FlourIn:             "100",
		HydrationIn:         "75",
		TangzhongPercentage: "10",
		TanghzhongRatio:     "1",
	}

	problems := models.FormErrors{}

	Calculator(&data, problems)

	if data.FlourOut != "95" {
		t.Errorf("Tangzhong Flour: Expected 95, Got: %s", data.FlourOut)
	}
	if data.HydrationOut != "70" {
		t.Errorf("Tangzhong Hydro: Expected: 70, Got: %s", data.HydrationOut)
	}
	if data.TangzhongFlour != "5" {
		t.Errorf("Tangzhong Flour: Expected 5, Got %s", data.TangzhongFlour)
	}
	if data.TangzhongHydration != "5" {
		t.Errorf("Tangzhong Hydro: Expected 5, Got: %s", data.TangzhongHydration)
	}
}

func TestTotalWeight(t *testing.T) {
	data := models.RecipeData{
		Calculator:    "bread",
		SubCalculator: "total-weight",
		DoughWeight:   "1200",
		EggIn:         "100",
		FatIn:         "100",
		HydrationIn:   "100",
		SaltIn:        "100",
		Leavener:      "Sourdough",
		SourdoughIn:   "100",
	}

	problems := models.FormErrors{}

	Calculator(&data, problems)

	if data.FlourIn != "200" {
		t.Errorf("Total Weight: Expected 200, Got: %s", data.FlourIn)
	}
}

func TestPanDimensions(t *testing.T) {
	tests := []struct {
		name        string
		shape       string
		measurement string
		height      string
		width       string
		depth       string
		diameter    string
		expected    string
	}{
		{"Square Centimeters", "square", "centimeters", "10", "10", "10", "10", "500"},
		{"Square Inches", "square", "inches", "10", "10", "10", "10", "8194"},
		{"Circle", "circle", "centimeters", "2", "", "", "100", "7854"},
	}

	problems := models.FormErrors{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := models.RecipeData{
				Calculator:    "bread",
				SubCalculator: "pan-dimension",
				Measurement:   tt.measurement,
				Shape:         tt.shape,
				Height:        tt.height,
				Width:         tt.width,
				Depth:         tt.depth,
				Diameter:      tt.diameter,
				HydrationIn:   "0",
				EggIn:         "0",
				FatIn:         "0",
				SugarIn:       "0",
				SaltIn:        "0",
			}

			Calculator(&data, problems)

			if data.FlourIn != tt.expected {
				t.Errorf("Got: %s, Expected: %s", data.FlourIn, tt.expected)
			}
		})
	}
}
