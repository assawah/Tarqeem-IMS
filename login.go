package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var LoginEnd string = "/login"

func login() {
	p := "login"
	E.GET(LoginEnd, func(c echo.Context) error {
		return c.Render(http.StatusOK, p, nil)
	})
}
