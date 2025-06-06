package main

import (
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

var templates = map[string]*template.Template{}
var pages = []string{"index", "about", "contact", "register", "login"}

func main() {
	loadPages()

	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", indexHandler)
	r.Get("/about", aboutHandler)
	r.Get("/contact", contactHandler)
	r.Get("/register", registerHandler)
	r.Get("/login", loginHandler)

	http.ListenAndServe(":80", r)
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
func contactHandler(w http.ResponseWriter, r *http.Request) {
	err := templates["contact"].Execute(w, nil)
	if err != nil {
		http.Error(w, "Template error", http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, w *http.Request) {
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
