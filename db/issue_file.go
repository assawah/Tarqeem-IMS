package db

import (
	"database/sql"
	"fmt"
)

type IssueFile struct {
	IssueID int
	FileID  int
}

// Create a new issue file
func CreateIssueFile(db *sql.DB, issueFile *IssueFile) error {
	query := `INSERT INTO issue_files (issue_id, file_id) VALUES (?, ?)`
	_, err := db.Exec(query, issueFile.IssueID, issueFile.FileID)
	return err
}

// Retrieve an issue file by IDs
func GetIssueFileByID(db *sql.DB, issueID, fileID int) (*IssueFile, error) {
	query := `SELECT issue_id, file_id FROM issue_files WHERE issue_id = ? AND file_id = ?`
	row := db.QueryRow(query, issueID, fileID)
	issueFile := &IssueFile{}
	err := row.Scan(&issueFile.IssueID, &issueFile.FileID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("issue file not found")
		}
		return nil, err
	}
	return issueFile, nil
}

// Delete an issue file by IDs
func DeleteIssueFile(db *sql.DB, issueID, fileID int) error {
	query := `DELETE FROM issue_files WHERE issue_id = ? AND file_id = ?`
	_, err := db.Exec(query, issueID, fileID)
	return err
}
