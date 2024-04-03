package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	// "github.com/tarqeem/ims/ent/user"

)

var LoginEnd string = "/login"
// var RegisterEnd string = "/register"



func login() {
	p := "login"
	E.GET(LoginEnd, func(c echo.Context) error {
		return c.Render(http.StatusOK, p, nil)
	})
	// E.POST(LoginEnd,func(c echo.Context) error {

	// 	l := &UserDTO{}

	// 	if err := c.Bind(l); err != nil {
	// 		return c.String(http.StatusBadRequest, "bad request")
	// 	}

	// 	u ,err := Client.User.Query().Where(user.EmailEQ(l.Email))
	// 	if err != nil {
	// 		E.Logger.Errorf("ent: %s", err.Error())
	// 		return c.Render(http.StatusBadRequest, "login", &UserDTO{Err: err.Error()})
	// 	}


	// 	sess, err := cookie(c)
	// 	if err != nil {
	// 		E.Logger.Errorf("ent: %s", err.Error())
	// 		return c.Render(http.StatusInternalServerError, "login", &UserDTO{Err: err.Error()})
	// 	}

	// 	// sess.Values["auth"] = "true"
	// 	sess.Values["id"] = u.ID

	// 	err = sess.Save(c.Request(), c.Response())
	// 	if err != nil {
	// 		E.Logger.Errorf("ent: %s", err.Error())
	// 		return c.Render(http.StatusInternalServerError, "login", &UserDTO{Err: err.Error()})
	// 	}

	// 	E.Logger.Infof("added user %+v", u)
	// 	return c.Redirect(http.StatusFound, DashboardEnd)	


	// })
}


