// commands/edit.go
package commands

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

func Edit(db *sql.DB, args []string) {

	if len(args) < 4 {
		fmt.Println("Usage: edit <id> <new task>")
		return
	}

	id, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Invalid: ID is not a number.")
	}

	Newtask := strings.Join(args[3:], " ")

	result, err := db.Exec("UPDATE tasks SET taskname = ? WHERE id = ?", Newtask, id)
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

	fmt.Printf("Task '%d' succesfully changed to '%s'.\n", id, Newtask)
}
