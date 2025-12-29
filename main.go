package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gwaDyckuL1/Ratio_Baking_Site/database"
	"github.com/gwaDyckuL1/Ratio_Baking_Site/handlers"
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
	_, err := database.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		log.Fatal(err)
	}
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.Handle("/", handlers.SessionMiddleware(database, handlers.IndexHandler(templates)))
	router.Handle("/about", handlers.SessionMiddleware(database, handlers.AboutHandler(templates)))
	router.Handle("/contact", handlers.SessionMiddleware(database, handlers.ContactHandler(templates)))
	router.Handle("/login", handlers.SessionMiddleware(database, handlers.LoginHandler(templates)))
	router.Handle("/loginSubmit", handlers.SessionMiddleware(database, handlers.LoginSubmitHandler(database)))
	router.Handle("/logout", handlers.SessionMiddleware(database, handlers.LogoutHandler(database)))
	router.Handle("/register", handlers.SessionMiddleware(database, handlers.RegisterHandler(templates)))
	router.Handle("/registrationSubmit", handlers.SessionMiddleware(database, handlers.RegistrationSubmitHandler(database)))

	router.Handle("/calculator/", handlers.SessionMiddleware(database, handlers.CalculatorIndexHandler(templates)))
	router.Handle("/calculator/bread", handlers.SessionMiddleware(database, handlers.CalculatorBreadHandler(templates)))
	router.Handle("/calculator/results", handlers.SessionMiddleware(database, handlers.CalcResultsHandler()))

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
