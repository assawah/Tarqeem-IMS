package db

import (
	"database/sql"
	"fmt"
)

type Issue struct {
	ID            int
	Title         string
	Description   string
	Creator       string
	Status        string
	Date          string
	ProjectIssues *int
}

// Create a new issue
func CreateIssue(db *sql.DB, issue *Issue) (*Issue, error) {
	query := `INSERT INTO issues (title, description, creator, status, date, project_issues) VALUES (?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, issue.Title, issue.Description, issue.Creator, issue.Status, issue.Date, issue.ProjectIssues)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	newIssue, err := GetIssueByID(db, int(id))
	if err != nil {
		return nil, err
	}
	return newIssue, nil
}

// Retrieve an issue by ID
func GetIssueByID(db *sql.DB, id int) (*Issue, error) {
	query := `SELECT id, title, description, creator, status, date, project_issues FROM issues WHERE id = ?`
	row := db.QueryRow(query, id)
	issue := &Issue{}
	err := row.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Creator, &issue.Status, &issue.Date, &issue.ProjectIssues)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("issue not found")
		}
		return nil, err
	}
	return issue, nil
}

// Update an issue's information
func UpdateIssue(db *sql.DB, issue *Issue) error {
	query := `UPDATE issues SET title = ?, description = ?, creator = ?, status = ?, date = ?, project_issues = ? WHERE id = ?`
	_, err := db.Exec(query, issue.Title, issue.Description, issue.Creator, issue.Status, issue.Date, issue.ProjectIssues, issue.ID)
	return err
}

// Delete an issue by ID
func DeleteIssue(db *sql.DB, id int) error {
	query := `DELETE FROM issues WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}

func GetIssuesByProjectID(db *sql.DB, projectID int) ([]*Issue, error) {
	query := `
        SELECT 
            id, 
            COALESCE(title, ''), 
            COALESCE(description, ''), 
            COALESCE(creator, ''), 
            COALESCE(status, ''), 
            COALESCE(date, ''), 
            project_issues 
        FROM issues 
        WHERE project_issues = ?
    `
	rows, err := db.Query(query, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve issues: %v", err)
	}
	defer rows.Close()

	var issues []*Issue
	for rows.Next() {
		var issue Issue
		err := rows.Scan(&issue.ID, &issue.Title, &issue.Description, &issue.Creator, &issue.Status, &issue.Date, &issue.ProjectIssues)
		if err != nil {
			return nil, fmt.Errorf("could not scan issue: %v", err)
		}
		issues = append(issues, &issue)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row error: %v", err)
	}

	return issues, nil
}
