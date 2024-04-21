package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tarqeem/ims/ent/user"
	. "github.com/tarqeem/ims/translate"
)

var RegisterEnd string = "/register"

type UserDTO struct {
	Name         string `form:"name"`
	Email        string `form:"email"`
	Title        string `form:"title"`
	Phone        string `form:"phone"`
	Password     string `form:"password"`
	Organization string `form:"organization"`
	Type         string `form:"type"`
	PageTitle    string
	Err          string
}

// TODO Fix regex values in tmplt
func register() {
	E.GET(RegisterEnd, func(c echo.Context) error {
		t := c.QueryParam("t")
		var m UserDTO
		m.Type = t

		if t == "c" {
			m.PageTitle = Message("regCoordinator")
		} else if t == "m" {
			m.PageTitle = Message("regMember")
		} else {
			return c.Render(http.StatusNotFound, "404", nil)
		}
		return c.Render(http.StatusOK, "register", m)
	})

	E.POST(RegisterEnd, func(c echo.Context) error {
		r := &UserDTO{}
		tmplt := "register"
		if err := c.Bind(r); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		var t user.Type
		if r.Type == "c" {
			t = user.TypeCoordinator
		} else if r.Type == "m" {
			t = user.TypeMember
		} else {
			err := errors.New("trying to add value: " + r.Type + " to type")
			E.Logger.Error(err)
			return c.Render(http.StatusBadRequest, tmplt, &UserDTO{Err: err.Error()})
		}

		u, err := Client.User.Create().
			SetName(r.Name).
			SetEmail(r.Email).
			SetTitle(r.Title).
			SetPhone(r.Phone).
			SetPassword(r.Password).
			SetOrganization(r.Organization).
			SetIsActive(true).
			SetType(t).Save(c.Request().Context())
		if err != nil {
			E.Logger.Errorf("ent: %s", err.Error())
			return c.Render(http.StatusBadRequest, tmplt, &UserDTO{Err: err.Error(), Type: r.Type})
		}

		sess, err := cookie(c)
		if err != nil {
			E.Logger.Errorf("ent: %s", err.Error())
			return c.Render(http.StatusInternalServerError, tmplt, &UserDTO{Err: err.Error(), Type: r.Type})
		}

		// sess.Values["auth"] = "true"
		sess.Values["id"] = u.ID

		err = sess.Save(c.Request(), c.Response())
		if err != nil {
			E.Logger.Errorf("ent: %s", err.Error())
			return c.Render(http.StatusInternalServerError, tmplt, &UserDTO{Err: err.Error(), Type: r.Type})
		}

		E.Logger.Infof("added user %+v", u)
		return c.Redirect(http.StatusFound, DashboardEnd)

	})
}
