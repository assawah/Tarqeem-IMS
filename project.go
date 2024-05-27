package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tarqeem/ims/db"
)

type ProjectDTO struct {
	Project *db.Project
	Issues  []*db.Issue
	Err     string
}

var ProjectEnd = "/project"

func projectView() {
	E.GET(ProjectEnd, func(c echo.Context) error {

		prName := c.QueryParam("name")
		data := ProjectDTO{}

		projectObject, err := db.GetProjectByName(DB, prName)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "fail", nil)
		}
		data.Project = projectObject

		issues, err := db.GetIssuesByProjectID(DB, projectObject.ID)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "fail", nil)
		}

		data.Issues = issues

		sess, err := cookie(c)
		if err != nil {
			E.Logger.Errorf("ent: %s", err.Error())
			return c.Render(http.StatusInternalServerError, "fail", nil)
		}

		// sess.Values["auth"] = "true"
		sess.Values["project_name"] = prName

		err = sess.Save(c.Request(), c.Response())
		if err != nil {
			E.Logger.Errorf("ent: %s", err.Error())
			return c.Render(http.StatusInternalServerError, "fail", nil)
		}

		return c.Render(http.StatusOK, "project", data)
	})
}
