package db

import (
	"database/sql"
	"fmt"
)

type IssueComment struct {
	IssueID   int
	CommentID int
}

// Create a new issue comment
func createIssueComment(db *sql.DB, issueComment *IssueComment) error {
	query := `INSERT INTO issue_comments (issue_id, comment_id) VALUES (?, ?)`
	_, err := db.Exec(query, issueComment.IssueID, issueComment.CommentID)
	return err
}

// Retrieve an issue comment by IDs
func getIssueCommentByID(db *sql.DB, issueID, commentID int) (*IssueComment, error) {
	query := `SELECT issue_id, comment_id FROM issue_comments WHERE issue_id = ? AND comment_id = ?`
	row := db.QueryRow(query, issueID, commentID)
	issueComment := &IssueComment{}
	err := row.Scan(&issueComment.IssueID, &issueComment.CommentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("issue comment not found")
		}
		return nil, err
	}
	return issueComment, nil
}

// Delete an issue comment by IDs
func deleteIssueComment(db *sql.DB, issueID, commentID int) error {
	query := `DELETE FROM issue_comments WHERE issue_id = ? AND comment_id = ?`
	_, err := db.Exec(query, issueID, commentID)
	return err
}
