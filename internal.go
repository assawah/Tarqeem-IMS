package main

import (
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func cookie(c echo.Context) (*sessions.Session, error) {
	sess, err := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 300,
		HttpOnly: true,
	}
	if err != nil {
		c.Logger().Errorf("couldn't get cookie; %v+", err)
	}
	return sess, err
}

func authenticated(c echo.Context) (id int, err error) {
	sess, err := cookie(c)
	if err != nil {
		return
	}
	if v, ok := sess.Values["id"]; ok {
		return v.(int), nil
	}
	return -1, UserNotFound
}

var UserNotFound error = fmt.Errorf("User was not found")
