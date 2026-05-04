package commands

import (
	"database/sql"
	"fmt"
	"strings"
)

func Add(db *sql.DB, args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: add <task>")
		return
	}

	taskName := strings.Join(args, " ")

	_, err := db.Exec(
		"INSERT INTO tasks (taskname, status) VALUES (?, ?)",
		taskName, 0,
	)
	if err != nil {
		fmt.Println("Error inserting task:", err)
		return
	}

	fmt.Println("Task added:", taskName)
}
