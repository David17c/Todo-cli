// commands/remove.go
package commands

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "modernc.org/sqlite"
)

func Remove(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: remove <ID>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid: ID is not a number.")
		return
	}

	db, err := sql.Open("sqlite", "file:tasks.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	result, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		fmt.Println("Query error:", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error checking result:", err)
		return
	}

	if rowsAffected == 0 {
		fmt.Println("No task found with that ID.")
		return
	}

	fmt.Printf("Task '%d' succesfully removed.", id)
}
