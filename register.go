package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/tarqeem/ims/db"
	"github.com/tarqeem/ims/translate"
)

var RegisterEnd string = "/register"

type UserDTO struct {
	Name         string `form:"name"`
	Email        string `form:"email"`
	Username     string `form:"username"`
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
			m.PageTitle = translate.Message("regCoordinator")
		} else if t == "m" {
			m.PageTitle = translate.Message("regMember")
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
		t := ""
		if r.Type == "c" {
			t = "coordinator"
		} else if r.Type == "m" {
			t = "member"
		}

		newUser := db.User{
			Name:         &r.Name,
			Password:     &r.Password,
			Email:        r.Email,
			Username:     r.Username,
			Phone:        &r.Phone,
			CreatedAt:    time.Now(),
			Organization: &r.Organization,
			IsActive:     true,
			Type:         t,
		}
		u, err := db.CreateUser(DB, newUser)
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
