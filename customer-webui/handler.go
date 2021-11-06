package main

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles("search.html"))

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/search", http.StatusFound)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "search.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
