package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tarqeem/ims/db"
)

// "github.com/tarqeem/ims/ent"

type DashboardDTO struct {
	User     *db.User
	Projects []*db.Project
	Err      string
}

var DashboardEnd string = "/dashboard"

func dashboard() {
	E.GET(DashboardEnd, func(c echo.Context) error {
		p := "dashboard"
		data := DashboardDTO{}
		id, err := authenticated(c)
		if err != nil && errors.Is(err, UserNotFound) {
			return c.Redirect(http.StatusTemporaryRedirect, LoginEnd)
		} else if err != nil {
			E.Logger.Errorf("authentication failed: %s", err.Error())
			return c.Render(http.StatusInternalServerError,
				"fail",
				&DashboardDTO{Err: err.Error()})
		}

		u, err := db.GetUserByID(DB, id)
		if err != nil {
			E.Logger.Error(err)
			return c.Render(http.StatusInternalServerError, "fail",
				&DashboardDTO{Err: err.Error()})
		}
		data.User = u
		memberProjects, err := db.GetProjectsByUserID(DB, u.ID)
		data.Projects = append(data.Projects, memberProjects...)

		coordinatorProjects, err := db.GetProjectsByCoordinatorID(DB, u.ID)
		if err != nil {
			return c.Render(http.StatusInternalServerError, "fail",
				&DashboardDTO{Err: err.Error()})
		}
		data.Projects = append(data.Projects, coordinatorProjects...)

		return c.Render(http.StatusOK, p, data)
	})

}
