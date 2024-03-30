package main

import (
	"context"
	"embed"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tarqeem/ims/ent"
	"github.com/tarqeem/ims/models"
	. "github.com/tarqeem/ims/translate"
	. "github.com/tarqeem/template/utl"
)

//go:embed pages/*
var views embed.FS

const debug = true

var executor TemplateExecutor

func main() {
	// Views
	Views = views
	TemplateFuncs = template.FuncMap{
		"message": Message,
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

	// Database
	client, err := ent.Open("sqlite3", "file:ent.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Controllers
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
		t := c.QueryParam("t")
		var m models.Register
		m.Type = "t"

		if t == "c" {
			m.PageTitle = Message("regCoordinator")
		} else if t == "m" {
			m.PageTitle = Message("regMember")
		} else {
			return c.Render(http.StatusNotFound, "404", nil)
		}
		return c.Render(http.StatusOK, "register", m)
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
