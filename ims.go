package main

import (
	"embed"
	"log"
	"net/http"
	"text/template"

	"github.com/tarqeem/template/utl"
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
	utl.TemplateFuncs = template.FuncMap{
		"message": func(k string) string {
			if val, ok := English[k]; ok {
				return val
			}
			m := "Couldn't find key " + k
			log.Println(m)
			return m

		},
	}

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

	// routes

	http.HandleFunc("/", NewHandler("login", nil))
	http.HandleFunc("/home", NewHandler("dashboard", nil))

	http.ListenAndServe(":8080", nil)
}
