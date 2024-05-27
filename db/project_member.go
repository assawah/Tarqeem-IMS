package db

import (
	"database/sql"
	"fmt"
)

type ProjectMember struct {
	ProjectID int
	UserID    int
}

// Create a new project member
func CreateProjectMember(db *sql.DB, projectMember *ProjectMember) error {
	query := `INSERT INTO project_members (project_id, user_id) VALUES (?, ?)`
	_, err := db.Exec(query, projectMember.ProjectID, projectMember.UserID)
	return err
}

// Retrieve a project member by IDs
func GetProjectMemberByID(db *sql.DB, projectID, userID int) (*ProjectMember, error) {
	query := `SELECT project_id, user_id FROM project_members WHERE project_id = ? AND user_id = ?`
	row := db.QueryRow(query, projectID, userID)
	projectMember := &ProjectMember{}
	err := row.Scan(&projectMember.ProjectID, &projectMember.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("project member not found")
		}
		return nil, err
	}
	return projectMember, nil
}

// Delete a project member by IDs
func DeleteProjectMember(db *sql.DB, projectID, userID int) error {
	query := `DELETE FROM project_members WHERE project_id = ? AND user_id = ?`
	_, err := db.Exec(query, projectID, userID)
	return err
}
