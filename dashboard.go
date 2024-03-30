package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tarqeem/ims/ent"
	"github.com/tarqeem/ims/ent/user"
)

type DashboardDTO struct {
	*ent.User
	Err string
}

var DashboardEnd string = "/dashboard"

func dashboard() {
	E.GET(DashboardEnd, func(c echo.Context) error {
		p := "dashboard"
		id, err := authenticated(c)
		if err != nil && errors.Is(err, UserNotFound) {
			return c.Redirect(http.StatusTemporaryRedirect, LoginEnd)
		} else if err != nil {
			E.Logger.Errorf("authentication failed: %s", err.Error())
			return c.Render(http.StatusInternalServerError,
				p,
				&DashboardDTO{Err: err.Error()})
		}

		u, err := Client.User.Query().Where(user.ID(id)).Only(c.Request().Context())
		if err != nil {
			E.Logger.Error(err)
			return c.Render(http.StatusInternalServerError, p,
				&DashboardDTO{Err: err.Error()})
		}

		return c.Render(http.StatusOK, p, &DashboardDTO{User: u})
	})

}
