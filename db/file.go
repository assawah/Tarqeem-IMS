package db

import (
	"database/sql"
	"fmt"
)

type File struct {
	ID       int
	FilePath *string
	FileName *string
	FileSize int
}

// Create a new file
func CreateFile(db *sql.DB, file *File) (*File, error) {
	query := `INSERT INTO files (file_path, file_name, file_size) VALUES (?, ?, ?)`
	result, err := db.Exec(query, file.FilePath, file.FileName, file.FileSize)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	newFile, err := GetFileByID(db, int(id))
	if err != nil {
		return nil, err
	}
	return newFile, nil
}

// Retrieve a file by ID
func GetFileByID(db *sql.DB, id int) (*File, error) {
	query := `SELECT id, COALESCE(file_path, ''), COALESCE(file_name, ''), COALESCE(file_size, 0) FROM files WHERE id = ?`
	row := db.QueryRow(query, id)
	file := &File{}
	err := row.Scan(&file.ID, &file.FilePath, &file.FileName, &file.FileSize)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("file not found")
		}
		return nil, err
	}
	return file, nil
}

// Update a file's information
func UpdateFile(db *sql.DB, file *File) error {
	query := `UPDATE files SET file_path = ?, file_name = ?, file_size = ? WHERE id = ?`
	_, err := db.Exec(query, file.FilePath, file.FileName, file.FileSize, file.ID)
	return err
}

// Delete a file by ID
func DeleteFile(db *sql.DB, id int) error {
	query := `DELETE FROM files WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}
