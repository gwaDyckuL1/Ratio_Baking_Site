package main

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	accounts "github.com/gwaDyckuL1/Ratio_Baking_Site/Accounts"
	"github.com/gwaDyckuL1/Ratio_Baking_Site/calculator"
	"github.com/gwaDyckuL1/Ratio_Baking_Site/database"
	"github.com/gwaDyckuL1/Ratio_Baking_Site/models"
	_ "github.com/mattn/go-sqlite3"
)

var pages = []string{"index", "about", "contact", "register", "login"}
var calcPages = []string{"calcIndex", "bread"}
var templates = map[string]*template.Template{}

func main() {
	loadPages()
	loadCalcPages()
	database := database.OpenDatabase()
	defer database.Close()

	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/about", aboutHandler)
	router.HandleFunc("/contact", contactHandler)
	router.HandleFunc("/login", loginHandler)
	router.HandleFunc("/register", registerHandler)
	router.HandleFunc("/registrationSubmit", registerationSubmitHandler(database))

	router.HandleFunc("/calculator/", calculatorIndexHandler)
	router.HandleFunc("/calculator/bread", breadCalcHandler)
	router.HandleFunc("/calculator/results", calcResultsHandler)

	server := http.Server{
		Addr:    ":80",
		Handler: router,
	}

	server.ListenAndServe()
}

func loadPages() {
	for _, page := range pages {
		tmpl := template.Must(template.ParseFiles(
			"templates/layout.html",
			"templates/"+page+".html",
		))
		templates[page] = tmpl
	}
}

func loadCalcPages() {
	for _, page := range calcPages {
		tmpl := template.Must(template.ParseFiles(
			"templates/layout.html",
			"templates/calculator/layout.html",
			"templates/calculator/"+page+".html",
		))
		templates[page] = tmpl
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["index"].Execute(w, nil)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["about"].Execute(w, nil)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

func breadCalcHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["bread"].Execute(w, nil)
	if err != nil {
		http.Error(w, "Template error with bread", http.StatusInternalServerError)
	}
}

func calculatorIndexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["calcIndex"].Execute(w, nil)
	if err != nil {
		http.Error(w, "Template error with Calculator", http.StatusInternalServerError)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["contact"].Execute(w, nil)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["login"].Execute(w, nil)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

func calcResultsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	data := models.RecipeData{
		Calculator:          r.FormValue("calculatorFor"),
		SubCalculator:       r.FormValue("calculator-bread"),
		Measurement:         r.FormValue("measurement"),
		Shape:               r.FormValue("shape"),
		Height:              r.FormValue("height"),
		Width:               r.FormValue("width"),
		Depth:               r.FormValue("depth"),
		Diameter:            r.FormValue("diameter"),
		FlourIn:             r.FormValue("flour"),
		DoughWeight:         r.FormValue("dough-weight"),
		HydrationIn:         r.FormValue("hydration"),
		EggIn:               r.FormValue("egg"),
		FatIn:               r.FormValue("fat"),
		SugarIn:             r.FormValue("sugar"),
		TangzhongPercentage: r.FormValue("tangzhong-percentage"),
		TanghzhongRatio:     r.FormValue("tangzhong-ratio"),
		SaltIn:              r.FormValue("salt"),
		Leavener:            r.FormValue("leavener-choice"),
		SourdoughIn:         r.FormValue("sourdough"),
		YeastIn:             r.FormValue("yeast"),
	}

	problems := models.FormErrors{}

	calculator.Calculator(&data, problems)

	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/calculator/layout.html",
		"templates/calculator/results.html",
	))
	tmpl.Execute(w, &data)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["register"].Execute(w, nil)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

func registerationSubmitHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		type Response struct {
			Ok      bool   `json:"ok"`
			Field   string `json:"field,omitempty"`
			Message string `json:"message,omitempty"`
		}

		err := r.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		data := models.RegistrationData{
			Username: r.FormValue("username"),
			Name:     r.FormValue("name"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("passwrod"),
		}

		emailUsed, err := accounts.CheckEmail(data.Email, db)
		if err != nil && err != sql.ErrNoRows {
			log.Printf("Database error checking email: %v", err)
			http.Error(w, "Internal Server Error. Please try again later", http.StatusInternalServerError)
			return
		}
		if emailUsed {
			if r.Header.Get("Accept") == "application/json" {
				json.NewEncoder(w).Encode(Response{
					Ok:      false,
					Field:   "email",
					Message: "This email already has an account.",
				})
			} else {
				tmpl := template.Must(template.ParseFiles(
					"templates/layout",
					"templates/register",
				))
				tmpl.Execute(w, map[string]string{
					"ErrorField": "email",
					"ErrorMsg":   "This email already has an account.",
				})
			}
			return
		}

		usernameUsed, err := accounts.CheckUserName(data.Username, db)
		if err != nil && err != sql.ErrNoRows {
			log.Printf("Database error checking username: %v", err)
			http.Error(w, "Internal Server Error. Please try again later", http.StatusInternalServerError)
			return
		}
		if usernameUsed {
			if r.Header.Get("Accept") == "application/json" {
				json.NewEncoder(w).Encode(Response{
					Ok:      false,
					Field:   "username",
					Message: "Username not available.  Please choose another.",
				})
			} else {
				tmpl := template.Must(template.ParseFiles(
					"templates/layout",
					"templates/register",
				))
				tmpl.Execute(w, map[string]string{
					"ErrorField": "username",
					"ErrorMsg":   "Username not available. Please choose another.",
				})
			}
			return
		}

		//Need to hash password still
		hashPassword := 0

		_, err = db.Exec(`INSERT INTO 
			users (username, name, email, password, role, create_date)
			VALUES (?, ?, ?, ?, ?, DATETIME("NOW"))`,
			data.Username, data.Name, data.Email, hashPassword, "User")

		if err != nil {
			log.Printf("Error inserting new user in database. Error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Response{Ok: false, Message: "Server error. Try again later."})
			return
		}

		json.NewEncoder(w).Encode(Response{Ok: true, Message: "Registration Successful"})
	}
}
