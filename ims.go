package main

import (
	"embed"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	. "github.com/tarqeem/ims/translate"
	. "github.com/tarqeem/template/utl"
)

//go:embed pages/*
var views embed.FS

const debug = true

var executor TemplateExecutor

func main() {
	Views = views
	TemplateFuncs = template.FuncMap{
		"message": func(k string) string {
			if val, ok := English[k]; ok {
				return val
			}
			m := "Couldn't find key " + k
			log.Println(m)
			return m

		},
	}
	ts, err := GetTemplates()

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

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("static"))
	e.Renderer = &etmplt{templates: ts}

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login", nil)
	})

	e.GET("/dashboard", func(c echo.Context) error {
		return c.Render(http.StatusOK, "dashboard", nil)
	})

	e.GET("/register", func(c echo.Context) error {
		return c.Render(http.StatusOK, "register", nil)
	})

	e.GET("/project", func(c echo.Context) error {
		return c.Render(http.StatusOK, "project", nil)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

type etmplt struct {
	templates *template.Template
}

func (t *etmplt) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if debug {
		return executor.ExecuteTemplate(w, name, data)
	}
	return t.templates.ExecuteTemplate(w, name, data)
}
