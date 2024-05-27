package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func registerIssueRoutes() {
	E.PUT("/updateItem/:id/:status", updateIssueStatus)
}

func updateIssueStatus(c echo.Context) error {
	// Parse request parameters
	itemID, _ := strconv.Atoi(c.Param("id"))
	newStatus := c.Param("status")

	fmt.Println(itemID)
	fmt.Println(newStatus)

	// Update item status in the database
	// You'll need to implement your database logic here
	// For example, using SQL:
	// _, err := db.Exec("UPDATE items SET status = ? WHERE id = ?", newStatus, itemID)

	// Handle errors
	// if err != nil {
	//     return err
	// }

	// Return success response
	return c.String(http.StatusOK, "Item status updated successfully")
}
