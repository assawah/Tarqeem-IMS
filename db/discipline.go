package db

import (
	"database/sql"
	"fmt"
)

type Discipline struct {
	ID   int
	Name string
}

// Create a new discipline
func createDiscipline(db *sql.DB, discipline *Discipline) error {
	query := `INSERT INTO disciplines (name) VALUES (?)`
	result, err := db.Exec(query, discipline.Name)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	discipline.ID = int(id)
	return nil
}

// Retrieve a discipline by ID
func getDisciplineByID(db *sql.DB, id int) (*Discipline, error) {
	query := `SELECT id, name FROM disciplines WHERE id = ?`
	row := db.QueryRow(query, id)
	discipline := &Discipline{}
	err := row.Scan(&discipline.ID, &discipline.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("discipline not found")
		}
		return nil, err
	}
	return discipline, nil
}

// Update a discipline's information
func updateDiscipline(db *sql.DB, discipline *Discipline) error {
	query := `UPDATE disciplines SET name = ? WHERE id = ?`
	_, err := db.Exec(query, discipline.Name, discipline.ID)
	return err
}

// Delete a discipline by ID
func deleteDiscipline(db *sql.DB, id int) error {
	query := `DELETE FROM disciplines WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}
