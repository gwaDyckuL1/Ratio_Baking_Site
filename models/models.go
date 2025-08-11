package models

type FormErrors map[string]string

type RecipeData struct {
	Calculator          string
	SubCalculator       string
	Measurement         string
	Shape               string
	Height              string
	Width               string
	Depth               string
	Radius              string
	Flour               string
	DoughWeight         string
	Hydration           string
	Fat                 string
	Sugar               string
	TangzhongPercentage string
	TangzhongFlour      string
	TangzhongHydration  string
	Salt                string
	Sourdough           string
	Yeast               string
}
