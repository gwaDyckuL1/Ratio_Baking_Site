package main

import (
	"html/template"
	"net/http"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/database"
	_ "github.com/mattn/go-sqlite3"
)

var pages = []string{"index", "about", "contact", "register", "login"}
var calcPages = []string{"calcIndex"}
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

	router.HandleFunc("/calculator/", calculatorIndexHandler)

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

func registerHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["register"].Execute(w, nil)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}
