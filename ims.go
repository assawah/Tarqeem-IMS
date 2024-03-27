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

var English map[string]string = map[string]string{
	"emailLoginNote": "Use the same email address you used for registeration.",
	"welcome":        "Welecome to Interface Management System (IMS).",
	"loginHelp":      "Please use your credentials to login or create a new account",
	"regCoordinator": "Register as a coordinator",
	"regMember":      "Register as a team member",
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	utl.Views = views
	utl.TemplateFuncs = template.FuncMap{
		"message": func(k string) string {
			return English[k]
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := executor.ExecuteTemplate(w, "base", nil)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error: "+err.Error(), 500)
		}

	})
	http.ListenAndServe(":8080", nil)
}
