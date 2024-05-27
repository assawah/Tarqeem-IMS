package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/tarqeem/ims/db"
)

var CreateProjectEnd = "/create-project"

type CreateProjectDTO struct {
	Name                string `form:"name"`
	Owner               string `form:"owner"`
	Location            string `form:"location"`
	Type                string `form:"type"`
	DeliveryStrategy    string `form:"deliveryStrategy"`
	CurrentState        string `form:"currentState"`
	ContractingStrategy string `form:"contractingStrategy"`
	DollarValue         int    `form:"dollarValue"`

	ExecutionLocation string `form:"executionLocation"`

	ProjectNature string   `form:"projectNature"`
	Leader        string   `form:"leader"`
	Members       []string `form:"members[]"`
	PageTitle     string
	Err           string
}

func createProject() {
	E.GET("/create-project", func(c echo.Context) error {
		return c.Render(http.StatusOK, "create-project", nil)
	})

	E.POST(CreateProjectEnd, func(c echo.Context) error {

		p := "success"

		r := &CreateProjectDTO{}
		if err := c.Bind(r); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		newProject := db.Project{
			Name:                  r.Name,
			Owner:                 r.Owner,
			Location:              r.Location,
			Type:                  r.Type,
			ProjectNature:         r.ProjectNature,
			DeliveryStrategies:    r.DeliveryStrategy,
			State:                 r.CurrentState,
			ContractingStrategies: r.ContractingStrategy,
			DollarValue:           r.DollarValue,
			ExecutionLocation:     r.ExecutionLocation,
		}

		proj, err := db.CreateProject(DB, &newProject)
		if err != nil {
			fmt.Print("Error creating project: " + err.Error())
			return c.Render(http.StatusInternalServerError, "fail",
				&CreateProjectDTO{Err: err.Error()})
		}
		err = AddCoordinator(c, proj)
		if err != nil {
			_ = db.DeleteProject(DB, proj.ID)
			fmt.Print("Error adding coordinator: " + err.Error())
			return c.Render(http.StatusInternalServerError, "fail",
				&CreateProjectDTO{Err: err.Error()})
		}
		err = AddLeader(proj, strings.TrimSpace(r.Leader))
		if err != nil {
			_ = db.DeleteProject(DB, proj.ID)
			fmt.Print("Error adding leader: " + err.Error())
			return c.Render(http.StatusInternalServerError, "fail",
				&CreateProjectDTO{Err: err.Error()})
		}
		r.Members = append(r.Members, strings.TrimSpace(r.Leader))
		err = AddMembers(proj, r.Members)
		if err != nil {
			_ = db.DeleteProject(DB, proj.ID)
			fmt.Print("Error adding members: " + err.Error())
			return c.Render(http.StatusInternalServerError, "fail",
				&CreateProjectDTO{Err: err.Error()})
		}

		err = c.Render(http.StatusOK, p, nil)
		if err != nil {
			return err
		}

		return c.Redirect(http.StatusOK, DashboardEnd)
	})
}

func AddCoordinator(c echo.Context, proj *db.Project) error {
	u, err := getCurrentUserID(c)
	if err != nil {
		return err
	}
	_, err = proj.AddCoordinator(DB, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func AddLeader(proj *db.Project, leader string) error {
	u, err := EnsureUser(leader)
	if err != nil {
		return err
	}
	_, err = proj.AddLeaderToProject(DB, u.ID)
	if err != nil {
		return err
	}
	return nil
}

func AddMembers(proj *db.Project, members []string) error {
	for _, member := range members {
		if member == "" {
			continue
		}
		trimmedMember := strings.TrimSpace(member)
		u, err := EnsureUser(trimmedMember)
		if err != nil {
			return err
		}
		_, err = proj.AddMemberToProject(DB, u.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func EnsureUser(email string) (*db.User, error) {
	u, err := db.GetUserByEmail(DB, strings.TrimSpace(email))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with email: %s", email)
		}
		return nil, err
	}
	return u, nil
}
