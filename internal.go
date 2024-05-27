package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/tarqeem/ims/db"
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
	fmt.Printf("+++++++++++++++++++++++++++context++++++++++++++++\n")
	fmt.Printf("%+v\n", c)
	fmt.Printf("+++++++++++++++++++++++++++context++++++++++++++++\n")
	sess, err := cookie(c)
	if err != nil {
		return
	}
	if v, ok := sess.Values["id"]; ok {
		fmt.Printf("+++++++++++++++++++++++++++id++++++++++++++++\n")
		fmt.Printf("%+v\n", v)
		fmt.Printf("+++++++++++++++++++++++++++id++++++++++++++++\n")
		return v.(int), nil
	}
	fmt.Printf("+++++++++++++++++++++++++++no id found++++++++++++++++\n")

	return -1, UserNotFound
}

func getCurrentUserID(c echo.Context) (*db.User, error) {
	fmt.Printf("+++++++++++++++++++++++++++context++++++++++++++++\n")
	fmt.Printf("%+v\n", c)
	fmt.Printf("+++++++++++++++++++++++++++context++++++++++++++++\n")

	id, err := authenticated(c)
	if err != nil {
		return nil, err
	}
	fmt.Printf("+++++++++++++++++++++++++++id++++++++++++++++\n")
	fmt.Printf("%+v\n", id)
	fmt.Printf("+++++++++++++++++++++++++++id++++++++++++++++\n")
	usr, err := db.GetUserByID(DB, id)
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
