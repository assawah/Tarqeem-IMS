package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tarqeem/ims/db"
)

var EditProjectEnd = "/edit-project"

type EditProjectDTO struct {
	ID       int    `form:"id"`
	Name     string `form:"name"`
	Owner    string `form:"owner"`
	Location string `form:"location"`
	Type     string `form:"type"`
	// DeliveryStrategies    string `form:"delivery_strategies"`
	// State                 string `form:"state"`
	// ContractingStrategies string `form:"contracting_strategies"`
	DeliveryStrategy    string `form:"deliveryStrategy"`
	CurrentState        string `form:"currentState"`
	ContractingStrategy string `form:"contractingStrategy"`
	DollarValue         int    `form:"dollarValue"`

	ExecutionLocation string   `form:"executionLocation"`
	ProjectNature     string   `form:"projectNature"`
	Leader            string   `form:"leader"`
	Members           []string `form:"members[]"`
	PageTitle         string
	Err               string
}

func editProject() {
	E.GET("/edit-project", func(c echo.Context) error {
		res := EditProjectDTO{}
		idstr := c.QueryParam("id")
		id, err := strconv.Atoi(idstr)
		if err != nil {
			return c.String(http.StatusBadRequest, "error in id")
		}
		project, err := db.GetProjectByID(DB, id)
		if err != nil {
			E.Logger.Error(err)
			return c.Render(http.StatusInternalServerError, "fail",
				nil)
		}

		projectLeader, err := db.GetProjectLeader(DB, project.ID)
		if err != nil {
			return c.String(http.StatusBadRequest, "error in GetProjectLeader")
		}

		projectMembers, err := db.GetProjectMembers(DB, project.ID)
		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusBadRequest, "error in GetProjectMembers")
		}
		res.ID = project.ID
		res.Name = project.Name
		res.Owner = project.Owner
		res.Location = project.Owner
		res.Type = string(project.Type)
		res.DeliveryStrategy = project.DeliveryStrategies
		res.CurrentState = project.State
		res.ContractingStrategy = project.ContractingStrategies
		res.DollarValue = project.DollarValue
		res.ProjectNature = project.ProjectNature
		res.Leader = projectLeader.Email
		ms := []string{}
		for _, v := range projectMembers {
			ms = append(ms, v.Email)
		}
		res.Members = ms

		return c.Render(http.StatusOK, "edit-project", res)
	})
	E.POST(EditProjectEnd, func(c echo.Context) error {

		r := &EditProjectDTO{}
		if err := c.Bind(r); err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		project := db.Project{
			ID:                    r.ID,
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

		err := db.UpdateProject(DB, &project)
		if err != nil {
			fmt.Print("Error editing project: " + err.Error())
			return c.Render(http.StatusInternalServerError, "fail",
				&EditProjectDTO{Err: err.Error()})
		}

		err = db.ClearProjectCoordinator(DB, project.ID)
		if err != nil {
			fmt.Print("Error clearing coordinator: " + err.Error())
			return c.Render(http.StatusInternalServerError, "fail",
				&EditProjectDTO{Err: err.Error()})
		}

		err = db.ClearProjectLeader(DB, project.ID)
		if err != nil {
			fmt.Print("Error clearing leader: " + err.Error())
			return c.Render(http.StatusInternalServerError, "fail",
				&EditProjectDTO{Err: err.Error()})
		}

		err = db.ClearProjectMembers(DB, project.ID)
		if err != nil {
			fmt.Print("Error clearing members: " + err.Error())
			return c.Render(http.StatusInternalServerError, "fail",
				&EditProjectDTO{Err: err.Error()})
		}

		err = AddCoordinator(c, &project)
		if err != nil {
			fmt.Print("Error adding coordinator: " + err.Error())
			return c.Render(http.StatusInternalServerError, "fail",
				&EditProjectDTO{Err: err.Error()})
		}
		err = AddLeader(&project, r.Leader)
		if err != nil {
			fmt.Print("Error adding leader: " + err.Error())
			return c.Render(http.StatusInternalServerError, "fail",
				&EditProjectDTO{Err: err.Error()})
		}
		r.Members = append(r.Members)
		err = AddMembers(&project, r.Members)
		if err != nil {
			fmt.Print("Error adding members: " + err.Error())
			return c.Render(http.StatusInternalServerError, "fail",
				&EditProjectDTO{Err: err.Error()})
		}

		err = c.Render(http.StatusOK, "editSuccess", nil)
		if err != nil {
			return err
		}

		return c.Redirect(http.StatusOK, DashboardEnd)

	})
}
