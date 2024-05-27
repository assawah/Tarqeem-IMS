package db

import (
	"database/sql"
	"fmt"
)

type UserLeaderOfProject struct {
	UserID    int
	ProjectID int
}

// Create a new user leader of project
func CreateUserLeaderOfProject(db *sql.DB, ulp *UserLeaderOfProject) error {
	query := `INSERT INTO user_leader_of_project (user_id, project_id) VALUES (?, ?)`
	_, err := db.Exec(query, ulp.UserID, ulp.ProjectID)
	return err
}

// Retrieve a user leader of project by IDs
func GetUserLeaderOfProjectByID(db *sql.DB, userID, projectID int) (*UserLeaderOfProject, error) {
	query := `SELECT user_id, project_id FROM user_leader_of_project WHERE user_id = ? AND project_id = ?`
	row := db.QueryRow(query, userID, projectID)
	ulp := &UserLeaderOfProject{}
	err := row.Scan(&ulp.UserID, &ulp.ProjectID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user leader of project not found")
		}
		return nil, err
	}
	return ulp, nil
}

// Delete a user leader of project by IDs
func DeleteUserLeaderOfProject(db *sql.DB, userID, projectID int) error {
	query := `DELETE FROM user_leader_of_project WHERE user_id = ? AND project_id = ?`
	_, err := db.Exec(query, userID, projectID)
	return err
}
