package db

import (
	"database/sql"
	"fmt"
)

type UserCoordinatorOfProject struct {
	UserID    int
	ProjectID int
}

// Create a new user coordinator of project
func CreateUserCoordinatorOfProject(db *sql.DB, ucp *UserCoordinatorOfProject) error {
	query := `INSERT INTO user_coordinator_of_project (user_id, project_id) VALUES (?, ?)`
	_, err := db.Exec(query, ucp.UserID, ucp.ProjectID)
	return err
}

// Retrieve a user coordinator of project by IDs
func GetUserCoordinatorOfProjectByProjectID(db *sql.DB, projectID int) (*UserCoordinatorOfProject, error) {
	query := `SELECT user_id, project_id FROM user_coordinator_of_project WHERE  project_id = ?`
	row := db.QueryRow(query, projectID)
	ucp := &UserCoordinatorOfProject{}
	err := row.Scan(&ucp.ProjectID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user coordinator of project not found")
		}
		return nil, err
	}
	return ucp, nil
}

// Delete a user coordinator of project by IDs
func DeleteUserCoordinatorOfProject(db *sql.DB, userID, projectID int) error {
	query := `DELETE FROM user_coordinator_of_project WHERE user_id = ? AND project_id = ?`
	_, err := db.Exec(query, userID, projectID)
	return err
}

func GetProjectsByCoordinatorID(db *sql.DB, userID int) ([]*Project, error) {
	query := `
        SELECT 
            p.id, 
            COALESCE(p.name, '') AS name, 
            COALESCE(p.owner, '') AS owner, 
            COALESCE(p.location, '') AS location, 
            COALESCE(p.type, '') AS type, 
            COALESCE(p.project_nature, '') AS project_nature, 
            COALESCE(p.delivery_strategies, '') AS delivery_strategies, 
            COALESCE(p.state, '') AS state, 
            COALESCE(p.contracting_strategies, '') AS contracting_strategies, 
            COALESCE(p.dollar_value, 0) AS dollar_value, 
            COALESCE(p.execution_location, '') AS execution_location
        FROM projects p
        JOIN user_coordinator_of_project ucp ON p.id = ucp.project_id
        WHERE ucp.user_id = ?
    `

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*Project
	for rows.Next() {
		var project Project
		if err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.Owner,
			&project.Location,
			&project.Type,
			&project.ProjectNature,
			&project.DeliveryStrategies,
			&project.State,
			&project.ContractingStrategies,
			&project.DollarValue,
			&project.ExecutionLocation); err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}
	return projects, err
}
