// commands/mark.go
package commands

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "modernc.org/sqlite"
)

func Mark(db *sql.DB, args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: mark <id>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid: ID is not a number.")
	}

	result, err := db.Exec("UPDATE tasks SET status = 1 WHERE id = ?", id)
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

	fmt.Printf("Task '%d' succesfully marked 'done'.", id)
}
