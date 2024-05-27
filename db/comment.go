package db

import (
	"database/sql"
	"fmt"
	"time"
)

type Comment struct {
	ID        int
	Content   string
	CreatedAt time.Time
}

// Create a new comment
func createComment(db *sql.DB, comment *Comment) error {
	query := `INSERT INTO comments (content, created_at) VALUES (?, ?)`
	result, err := db.Exec(query, comment.Content, comment.CreatedAt)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	comment.ID = int(id)
	return nil
}

// Retrieve a comment by ID
func getCommentByID(db *sql.DB, id int) (*Comment, error) {
	query := `SELECT id, content, COALESCE(created_at, CURRENT_TIMESTAMP) FROM comments WHERE id = ?`
	row := db.QueryRow(query, id)
	comment := &Comment{}
	err := row.Scan(&comment.ID, &comment.Content, &comment.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("comment not found")
		}
		return nil, err
	}
	return comment, nil
}

// Update a comment's information
func updateComment(db *sql.DB, comment *Comment) error {
	query := `UPDATE comments SET content = ?, created_at = ? WHERE id = ?`
	_, err := db.Exec(query, comment.Content, comment.CreatedAt, comment.ID)
	return err
}

// Delete a comment by ID
func deleteComment(db *sql.DB, id int) error {
	query := `DELETE FROM comments WHERE id = ?`
	_, err := db.Exec(query, id)
	return err
}
