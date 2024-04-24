package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/tarqeem/ims/ent/project"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// Issue represents the structure of an issue
type Issue struct {
	Title       string  `form:"title"`
	Description string  `form:"description"`
	Files       []*File `form:"files"`
}

// File represents a file uploaded with an issue
type File struct {
	Name     string
	Size     int64
	TempPath string
}

type CreateIssueResponse struct {
	Message string `json:"message"`
}

func createIssue() {
	E.POST("/create-issue", func(c echo.Context) error {
		// Parse the form data

		issue, err := parseIssue(c)

		usr, err := getCurrentUserID(c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to create issue",
			})
		}

		projName, err := getProjectName(c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to create issue",
			})
		}
		projObject, err := Client.Project.Query().
			Where(project.NameEQ(projName)).
			Only(context.Background())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to create issue",
			})
		}

		issueObject, err := Client.Issue.Create().
			SetTitle(issue.Title).
			SetDescription(issue.Description).
			SetProject(projObject).
			SetCreator(usr.Username).
			Save(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Failed to create issue",
			})
		}

		for _, file := range issue.Files {
			_, err = Client.File.Create().
				SetFileName(file.Name).
				SetFilePath(file.TempPath).
				SetFileSize(file.Size).
				AddIssue(issueObject).
				Save(c.Request().Context())
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "Failed to create issue",
				})
			}
		}

		// Return a success response
		return c.JSON(http.StatusOK, &CreateIssueResponse{
			Message: "Issue Created Successfully",
		})
	})
}

func parseIssue(c echo.Context) (*Issue, error) {
	form, err := c.MultipartForm()
	if err != nil {
		return nil, err
	}
	defer form.RemoveAll()

	// Parse issue details from the form
	title := c.FormValue("title")
	description := c.FormValue("description")

	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		return nil, err
	}

	// Parse uploaded files
	var files []*File
	for _, headers := range form.File {
		for _, header := range headers {
			// Create a temporary file to save the uploaded file

			file, err := header.Open()
			if err != nil {
				return nil, err
			}
			defer file.Close()

			// Create a temporary file to save the uploaded file
			tempPath := filepath.Join("./uploads", header.Filename)
			outFile, err := os.Create(tempPath)
			if err != nil {
				return nil, err
			}
			defer outFile.Close()

			// Copy the contents of the uploaded file to the temporary file
			if _, err = io.Copy(outFile, file); err != nil {
				return nil, err
			}

			// Add the uploaded file to the list of files
			files = append(files, &File{
				Name:     header.Filename,
				Size:     header.Size,
				TempPath: tempPath,
			})
		}
	}

	// Create an issue object with the parsed data
	return &Issue{
		Title:       title,
		Description: description,
		Files:       files,
	}, nil

}
