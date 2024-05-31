package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

type Comment struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	TaskID    int       `json:"task_id"`
	UserID    int       `json:"user_id"`
}

func createComments() {
    E.GET("/comments", func(c echo.Context) error {
        return c.Render(http.StatusOK, "comments", nil)
    })
}

func getComments(c echo.Context) error {
	// taskID := r.URL.Query().Get("task_id")
	// if taskID == "" {
	//     http.Error(w, "task_id is required", http.StatusBadRequest)
	//     return
	// }

	// db, err := sql.Open("sqlite3", "./project_management.db")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer db.Close()

	// rows, err := db.Query("SELECT id, content, created_at, task_id, user_id FROM comments WHERE task_id = ? ORDER BY created_at", taskID)
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer rows.Close()

	// var comments []Comment
	// for rows.Next() {
	//     var comment Comment
	//     if err := rows.Scan(&comment.ID, &comment.Content, &comment.CreatedAt, &comment.TaskID, &comment.UserID); err != nil {
	//         log.Fatal(err)
	//     }
	//     comments = append(comments, comment)
	// }
	comments := []Comment{Comment{
		ID:        1,
		Content:   "comment1 comment1 comment1 comment1 comment1 comment1 comment1 comment1",
		CreatedAt: time.Now(),
		TaskID:    1,
		UserID:    1,
	}, Comment{
		ID:        2,
		Content:   "comment2 comment2 comment2 comment2 comment2 comment2 comment2 comment2",
		CreatedAt: time.Now(),
		TaskID:    1,
		UserID:    2,
	}, Comment{
		ID:        3,
		Content:   "comment3 comment3 comment3 comment3 comment3 comment3 comment3 comment3",
		CreatedAt: time.Now(),
		TaskID:    1,
		UserID:    3,
	}, Comment{
		ID:        4,
		Content:   "comment4 comment4 comment4 comment4 comment4 comment4 comment4 comment4",
		CreatedAt: time.Now(),
		TaskID:    1,
		UserID:    4,
	}}

	return c.JSON(http.StatusOK, comments)
}


