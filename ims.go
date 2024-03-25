package main

import (
	"embed"
	"github.com/tarqeem/ims/utl"
	"log"
	"net/http"
)

//go:embed public/*
var public embed.FS

//go:embed pages/*
var views embed.FS

const debug = true

var executor utl.TemplateExecutor

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	utl.Views = views

	ts, err := utl.GetTemplates()

	if err != nil {
		log.Fatal(err)
	}
	if debug {

		if err != nil {
			log.Fatal(err)
		}
		executor = utl.DebugTemplateExecutor{}

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
