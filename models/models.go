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
	Diameter            string
	FlourIn             string
	FlourOut            string
	DoughWeight         string
	HydrationIn         string
	HydrationOut        string
	FatIn               string
	FatOut              string
	SugarIn             string
	SugarOut            string
	TangzhongPercentage string
	TanghzhongRatio     string
	TangzhongFlour      string
	TangzhongHydration  string
	SaltIn              string
	SaltOut             string
	Leavener            string
	LeavenerOut         string
	SourdoughIn         string
	YeastIn             string
}
