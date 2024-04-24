package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/tarqeem/ims/ent"
	"github.com/tarqeem/ims/ent/user"
	"math/rand"
	"time"
)

var UserNotFound error = fmt.Errorf("User was not found")

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

func getCurrentUserID(c echo.Context) (*ent.User, error) {
	id, err := authenticated(c)
	if err != nil {
		return nil, err
	}
	usr, err := Client.User.Query().Where(user.ID(id)).Only(c.Request().Context())
	if err != nil {
		E.Logger.Error(err)
		return nil, err
	}
	return usr, nil
}

func getProjectName(c echo.Context) (name string, err error) {
	sess, err := cookie(c)
	if err != nil {
		return
	}
	if v, ok := sess.Values["project_name"]; ok {
		return v.(string), nil
	}
	return "", fmt.Errorf("project Name not found")
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
