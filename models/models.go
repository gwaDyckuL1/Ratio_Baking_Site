package models

type FormErrors map[string]string

type Login struct {
	Username string
	Password string
}

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
	EggIn               string
	EggGrams            string
	EggWhole            string
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
	Notes               string
}

type RegistrationData struct {
	Username string
	Name     string
	Email    string
	Password string
}

type Response struct {
	Ok      bool   `json:"ok"`
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

type Session struct {
	LoggedIn     bool
	Username     string
	Name         string
	UserID       int
	SessionToken string
}

type WebData struct {
	RecipeData *RecipeData
	Session    *Session
	RecipeJSON string
}
