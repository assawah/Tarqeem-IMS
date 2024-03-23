package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed pages/*
var views embed.FS

func getTemplates() (*template.Template, error) {
	files, err := getFSFilesRecursively(&views, "pages")
	if err != nil {
		log.Fatal(err)
	}

	for _, value := range files {
		fmt.Println(value)
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return ts, err

}

//go:embed public/*
var public embed.FS

func main() {
	ts, err := getTemplates()
	if err != nil {
		log.Fatal(err)
	}

	staticHandler := http.FileServer(http.FS(public))
	http.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", 500)
		}

	})
	http.ListenAndServe(":8080", nil)
}
