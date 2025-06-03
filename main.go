package main

import (
	"net/http"
	"text/template"
)

func main() {

	http.HandleFunc("/", indexHandle)

	http.ListenAndServe(":80", nil)
}

func indexHandle(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	tmpl.Execute(w, nil)
}
