package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed public/*
var public embed.FS

const debug = true

var executor TemplateExecutor

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	ts, err := getTemplates()

	if err != nil {
		log.Fatal(err)
	}
	if debug {

		if err != nil {
			log.Fatal(err)
		}
		executor = DebugTemplateExecutor{}

	} else {
		executor = ts
	}

	staticHandler := http.FileServer(http.FS(public))
	http.Handle("/static/", http.StripPrefix("/static/", staticHandler))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := executor.ExecuteTemplate(w, "base", nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error: "+err.Error(), 500)
		}

	})
	http.ListenAndServe(":8080", nil)
}
