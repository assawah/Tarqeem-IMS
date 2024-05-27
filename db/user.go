package db

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID           int
	Name         *string
	Password     *string
	Email        string
	Username     string
	Phone        *string
	CreatedAt    time.Time
	Organization *string
	IsActive     bool
	Type         string
}

// Create a new user
func CreateUser(db *sql.DB, user User) (*User, error) {
	query := `INSERT INTO users (name, password, email, username, phone, created_at, organization, is_active, type) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := db.Exec(query, user.Name, user.Password, user.Email, user.Username, user.Phone, user.CreatedAt, user.Organization, user.IsActive, user.Type)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	u, err := GetUserByID(db, int(id))
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Retrieve a user by ID
func GetUserByID(db *sql.DB, id int) (*User, error) {
	query := `SELECT id, COALESCE(name, ''), COALESCE(password, ''), email, username, COALESCE(phone, ''), COALESCE(created_at, CURRENT_TIMESTAMP), COALESCE(organization, ''), is_active, type FROM users WHERE id = ?`
	row := db.QueryRow(query, id)
	user := &User{}
	var createdAtStr string
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Username, &user.Phone, &createdAtStr, &user.Organization, &user.IsActive, &user.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	if createdAtStr != "" {
		createdAt, err := time.Parse("2006-01-02 15:04:05.999999-07:00", createdAtStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse created_at: %v", err)
		}
		user.CreatedAt = createdAt
	}
	return user, nil
}

// Update a user's information
func UpdateUser(db *sql.DB, user *User) error {
	query := `UPDATE users SET name = ?, password = ?, email = ?, username = ?, phone = ?, created_at = ?, organization = ?, is_active = ?, type = ? WHERE id = ?`
	_, err := db.Exec(query, user.Name, user.Password, user.Email, user.Username, user.Phone, user.CreatedAt, user.Organization, user.IsActive, user.Type, user.ID)
	return err
}

// Delete a user by ID
func DeleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}

// GetUserByEmail retrieves a user by email from the users table
func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	query := `
        SELECT 
            id, 
            COALESCE(name, ''), 
            COALESCE(password, ''), 
            COALESCE(email, ''), 
            COALESCE(username, ''), 
            COALESCE(phone, ''), 
            COALESCE(created_at, ''), 
            COALESCE(organization, ''), 
            COALESCE(is_active, 0), 
            COALESCE(type, '') 
        FROM users 
        WHERE email = ?
    `
	fmt.Printf("=================emailDB++++++++++++++++++=\n")
	fmt.Printf("1%+v1\n", email)
	fmt.Printf("=================emailDB++++++++++++++++++=\n")
	row := db.QueryRow(query, email)
	var createdAtStr string
	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Username, &user.Phone, &createdAtStr, &user.Organization, &user.IsActive, &user.Type)
	fmt.Printf("=================errDB++++++++++++++++++=\n")
	fmt.Printf("1%+v1\n", err)
	fmt.Printf("=================errDB++++++++++++++++++=\n")
	if err != nil {
		return nil, err
	}
	if createdAtStr != "" {
		createdAt, err := time.Parse("2006-01-02 15:04:05.999999-07:00", createdAtStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse created_at: %v", err)
		}
		user.CreatedAt = createdAt
	}
	return &user, nil
}

// GetProjectLeader retrieves the leader of a specific project by project ID
func GetProjectLeader(db *sql.DB, projectID int) (*User, error) {
	query := `
        SELECT 
            u.id, 
            COALESCE(u.name, ''), 
            COALESCE(u.password, ''), 
            COALESCE(u.email, ''), 
            COALESCE(u.username, ''), 
            COALESCE(u.phone, ''), 
            COALESCE(u.created_at, ''), 
            COALESCE(u.organization, ''), 
            COALESCE(u.is_active, 0), 
            COALESCE(u.type, '') 
        FROM user_leader_of_project ulp
        JOIN users u ON ulp.user_id = u.id
        WHERE ulp.project_id = ?
    `
	row := db.QueryRow(query, projectID)

	var user User
	var createdAtStr string
	err := row.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Username, &user.Phone, &createdAtStr, &user.Organization, &user.IsActive, &user.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no leader found for project ID: %d", projectID)
		}
		return nil, err
	}
	if createdAtStr != "" {
		createdAt, err := time.Parse("2006-01-02 15:04:05.999999-07:00", createdAtStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse created_at: %v", err)
		}
		user.CreatedAt = createdAt
	}
	return &user, nil
}

// GetProjectMembers retrieves the members of a specific project by project ID
func GetProjectMembers(db *sql.DB, projectID int) ([]*User, error) {
	query := `
        SELECT 
            u.id, 
            COALESCE(u.name, ''), 
            COALESCE(u.password, ''), 
            COALESCE(u.email, ''), 
            COALESCE(u.username, ''), 
            COALESCE(u.phone, ''), 
            COALESCE(u.created_at, ''), 
            COALESCE(u.organization, ''), 
            COALESCE(u.is_active, 0), 
            COALESCE(u.type, '') 
        FROM project_members pm
        JOIN users u ON pm.user_id = u.id
        WHERE pm.project_id = ?
    `
	rows, err := db.Query(query, projectID)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve project members: %v", err)
	}
	defer rows.Close()

	var members []*User
	for rows.Next() {
		var user User
		var createdAtStr string
		err := rows.Scan(&user.ID, &user.Name, &user.Password, &user.Email, &user.Username, &user.Phone, &createdAtStr, &user.Organization, &user.IsActive, &user.Type)
		if err != nil {
			return nil, fmt.Errorf("could not scan user: %v", err)
		}
		members = append(members, &user)
		if createdAtStr != "" {
			createdAt, err := time.Parse("2006-01-02 15:04:05.999999-07:00", createdAtStr)
			if err != nil {
				return nil, fmt.Errorf("failed to parse created_at: %v", err)
			}
			user.CreatedAt = createdAt
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row error: %v", err)
	}

	return members, nil
}
