package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

var interfacePointEnd = "/interface-point"

// type CreateProjectDTO struct {
// 	Name                string `form:"name"`
// 	Owner               string `form:"owner"`
// 	Location            string `form:"location"`
// 	Type                string `form:"type"`
// 	DeliveryStrategy    string `form:"deliveryStrategy"`
// 	CurrentState        string `form:"currentState"`
// 	ContractingStrategy string `form:"contractingStrategy"`
// 	DollarValue         int    `form:"dollarValue"`

// 	ExecutionLocation string `form:"executionLocation"`

// 	ProjectNature string   `form:"projectNature"`
// 	Leader        string   `form:"leader"`
// 	Members       []string `form:"members[]"`
// 	PageTitle     string
// 	Err           string
// }

func interfacePoint() {
	E.GET(interfacePointEnd, func(c echo.Context) error {
		return c.Render(http.StatusOK, "interface-point", nil)
	})

	// E.POST(CreateProjectEnd, func(c echo.Context) error {

	// 	p := "success"

	// 	r := &CreateProjectDTO{}
	// 	if err := c.Bind(r); err != nil {
	// 		return c.String(http.StatusBadRequest, "bad request")
	// 	}

	// 	newProject := db.Project{
	// 		Name:                  r.Name,
	// 		Owner:                 r.Owner,
	// 		Location:              r.Location,
	// 		Type:                  r.Type,
	// 		ProjectNature:         r.ProjectNature,
	// 		DeliveryStrategies:    r.DeliveryStrategy,
	// 		State:                 r.CurrentState,
	// 		ContractingStrategies: r.ContractingStrategy,
	// 		DollarValue:           r.DollarValue,
	// 		ExecutionLocation:     r.ExecutionLocation,
	// 	}

	// 	proj, err := db.CreateProject(DB, &newProject)
	// 	if err != nil {
	// 		fmt.Print("Error creating project: " + err.Error())
	// 		return c.Render(http.StatusInternalServerError, "fail",
	// 			&CreateProjectDTO{Err: err.Error()})
	// 	}
	// 	err = AddCoordinator(c, proj)
	// 	if err != nil {
	// 		_ = db.DeleteProject(DB, proj.ID)
	// 		fmt.Print("Error adding coordinator: " + err.Error())
	// 		return c.Render(http.StatusInternalServerError, "fail",
	// 			&CreateProjectDTO{Err: err.Error()})
	// 	}
	// 	err = AddLeader(proj, strings.TrimSpace(r.Leader))
	// 	if err != nil {
	// 		_ = db.DeleteProject(DB, proj.ID)
	// 		fmt.Print("Error adding leader: " + err.Error())
	// 		return c.Render(http.StatusInternalServerError, "fail",
	// 			&CreateProjectDTO{Err: err.Error()})
	// 	}
	// 	r.Members = append(r.Members, strings.TrimSpace(r.Leader))
	// 	err = AddMembers(proj, r.Members)
	// 	if err != nil {
	// 		_ = db.DeleteProject(DB, proj.ID)
	// 		fmt.Print("Error adding members: " + err.Error())
	// 		return c.Render(http.StatusInternalServerError, "fail",
	// 			&CreateProjectDTO{Err: err.Error()})
	// 	}

	// 	err = c.Render(http.StatusOK, p, nil)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	return c.Redirect(http.StatusOK, DashboardEnd)
	// })
}

type Option struct {
    Value string `json:"value"`
    Text  string `json:"text"`
}

func getOptions(c echo.Context) error {
    options := []Option{
        {Value: "option1", Text: "Oprion1"},
        {Value: "option2", Text: "Option2"},
        {Value: "option3", Text: "Option3"},
    }
    return c.JSON(http.StatusOK, options)
}

func getFiles(c echo.Context) error {
    // taskID := c.QueryParam("task_id")
    // Fetch files associated with the taskID from the database or filesystem
    // Here, we assume a simple list of files for demonstration
    files := []map[string]string{
        {"filePath": "/uploads/file1.pdf", "fileName": "file1.pdf"},
        {"filePath": "/uploads/file2.pdf", "fileName": "file2.pdf"},
    }
    return c.JSON(http.StatusOK, files)
}