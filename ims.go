package main

import (
	"html/template" // New import
	"log"           // New import
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Include the navigation partial in the template files.
	files := []string{
		"./pages/base.tmpl",
		"./pages/partials/nav.tmpl",
		"./pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
