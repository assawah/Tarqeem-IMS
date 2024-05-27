package db

import (
	"database/sql"
	"fmt"
)

type Project struct {
	ID                    int
	Name                  string
	Owner                 string
	Location              string
	Type                  string
	ProjectNature         string
	DeliveryStrategies    string
	State                 string
	ContractingStrategies string
	DollarValue           int
	ExecutionLocation     string
}

// Create a new project
func CreateProject(db *sql.DB, project *Project) (*Project, error) {
	query := `INSERT INTO projects (name, owner, location, type, project_nature, delivery_strategies, state, contracting_strategies, dollar_value, execution_location) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, project.Name, project.Owner, project.Location, project.Type, project.ProjectNature, project.DeliveryStrategies, project.State, project.ContractingStrategies, project.DollarValue, project.ExecutionLocation)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	newProject, err := GetProjectByID(db, int(id))
	if err != nil {
		return nil, err
	}
	return newProject, nil
}

// Retrieve a project by ID
func GetProjectByID(db *sql.DB, id int) (*Project, error) {
	query := `SELECT id, name, owner, location, type, project_nature, delivery_strategies, state, contracting_strategies, dollar_value, execution_location FROM projects WHERE id = ?`
	row := db.QueryRow(query, id)
	project := &Project{}
	err := row.Scan(&project.ID, &project.Name, &project.Owner, &project.Location, &project.Type, &project.ProjectNature, &project.DeliveryStrategies, &project.State, &project.ContractingStrategies, &project.DollarValue, &project.ExecutionLocation)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("project not found")
		}
		return nil, err
	}
	return project, nil
}

func GetProjectByName(db *sql.DB, name string) (*Project, error) {
	query := `SELECT id, name, owner, location, type, project_nature, delivery_strategies, state, contracting_strategies, dollar_value, execution_location FROM projects WHERE name = ?`
	row := db.QueryRow(query, name)
	project := &Project{}
	err := row.Scan(&project.ID, &project.Name, &project.Owner, &project.Location, &project.Type, &project.ProjectNature, &project.DeliveryStrategies, &project.State, &project.ContractingStrategies, &project.DollarValue, &project.ExecutionLocation)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("project not found")
		}
		return nil, err
	}
	return project, nil
}

// Update a project's information
func UpdateProject(db *sql.DB, project *Project) error {
	query := `UPDATE projects SET name = ?, owner = ?, location = ?, type = ?, project_nature = ?, delivery_strategies = ?, state = ?, contracting_strategies = ?, dollar_value = ?, execution_location = ? WHERE id = ?`
	_, err := db.Exec(query, project.Name, project.Owner, project.Location, project.Type, project.ProjectNature, project.DeliveryStrategies, project.State, project.ContractingStrategies, project.DollarValue, project.ExecutionLocation, project.ID)
	return err
}

// Delete a project by ID
func DeleteProject(db *sql.DB, id int) error {
	query := `DELETE FROM projects WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}

// GetProjectsByUserID returns all projects that a user is a member of
func GetProjectsByUserID(db *sql.DB, userID int) ([]*Project, error) {
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
        JOIN project_members pm ON p.id = pm.project_id
        WHERE pm.user_id = ?
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func (p *Project) AddCoordinator(db *sql.DB, userId int) (*UserCoordinatorOfProject, error) {
	ucp := UserCoordinatorOfProject{
		UserID:    userId,
		ProjectID: p.ID,
	}
	err := CreateUserCoordinatorOfProject(db, &ucp)
	if err != nil {
		return nil, err
	}
	return &ucp, err

}

func (p *Project) AddMemberToProject(db *sql.DB, userID int) (*ProjectMember, error) {
	pm := ProjectMember{
		ProjectID: p.ID,
		UserID:    userID,
	}
	err := CreateProjectMember(db, &pm)
	if err != nil {
		return nil, err
	}
	return &pm, nil
}

// AddLeaderToProject adds a leader to a project by inserting a record into the user_leader_of_project table
func (p *Project) AddLeaderToProject(db *sql.DB, userID int) (*UserLeaderOfProject, error) {
	ulp := UserLeaderOfProject{
		UserID:    userID,
		ProjectID: p.ID,
	}
	err := CreateUserLeaderOfProject(db, &ulp)
	if err != nil {
		return nil, err
	}
	return &ulp, nil
}

// ClearProjectCoordinator removes the coordinator from a specific project by project ID
func ClearProjectCoordinator(db *sql.DB, projectID int) error {
	query := "DELETE FROM user_coordinator_of_project WHERE project_id = ?"
	_, err := db.Exec(query, projectID)
	if err != nil {
		return fmt.Errorf("could not clear project coordinator: %v", err)
	}
	return nil
}

// ClearProjectLeader removes the leader from a specific project by project ID
func ClearProjectLeader(db *sql.DB, projectID int) error {
	query := "DELETE FROM user_leader_of_project WHERE project_id = ?"
	_, err := db.Exec(query, projectID)
	if err != nil {
		return fmt.Errorf("could not clear project leader: %v", err)
	}
	return nil
}

// ClearProjectMembers removes all members from a specific project by project ID
func ClearProjectMembers(db *sql.DB, projectID int) error {
	query := "DELETE FROM project_members WHERE project_id = ?"
	_, err := db.Exec(query, projectID)
	if err != nil {
		return fmt.Errorf("could not clear project members: %v", err)
	}
	return nil
}
