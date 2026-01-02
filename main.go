package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db := database.OpenDatabase()
	defer db.Close()
	_, err := db.Exec("PRAGMA journal_mode=WAL;")
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()

	database.SessionCleanUp(ticker.C, ctx, database.DeleteOldSessions(db))

	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	router.Handle("/", handlers.SessionMiddleware(db, handlers.IndexHandler(templates)))
	router.Handle("/about", handlers.SessionMiddleware(db, handlers.AboutHandler(templates)))
	router.Handle("/contact", handlers.SessionMiddleware(db, handlers.ContactHandler(templates)))
	router.Handle("/login", handlers.SessionMiddleware(db, handlers.LoginHandler(templates)))
	router.Handle("/loginSubmit", handlers.SessionMiddleware(db, handlers.LoginSubmitHandler(db)))
	router.Handle("/logout", handlers.SessionMiddleware(db, handlers.LogoutHandler(db)))
	router.Handle("/register", handlers.SessionMiddleware(db, handlers.RegisterHandler(templates)))
	router.Handle("/registrationSubmit", handlers.SessionMiddleware(db, handlers.RegistrationSubmitHandler(db)))

	router.Handle("/saveRecipe", handlers.SessionMiddleware(db, handlers.SaveRecipeHandler(db)))

	router.Handle("/calculator/", handlers.SessionMiddleware(db, handlers.CalculatorIndexHandler(templates)))
	router.Handle("/calculator/bread", handlers.SessionMiddleware(db, handlers.CalculatorBreadHandler(templates)))
	router.Handle("/calculator/results", handlers.SessionMiddleware(db, handlers.CalcResultsHandler()))

	server := http.Server{
		Addr:    ":80",
		Handler: router,
	}

	go func() {
		log.Println("Starting server on: ", server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalln("Server failed. ", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Println("Shutting server down...")

	cancel()

	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	err = server.Shutdown(ctxShutdown)
	if err != nil {
		log.Fatalln("Forced shutdown: ", err)
	}
	log.Println("Server shutdown, everything is clean!")
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
