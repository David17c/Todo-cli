// commands/add.go
package commands

import (
	"database/sql"
	"fmt"
	"strings"

	_ "modernc.org/sqlite"
)

func Add(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: add <taskname>")
		return
	}

	taskName := strings.Join(args, " ")

	db, err := sql.Open("sqlite", "file:tasks.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY,
		taskname TEXT NOT NULL,
		status INTEGER NOT NULL
	);
	`)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	_, err = db.Exec(
		"INSERT INTO tasks (taskname, status) VALUES (?, ?)",
		taskName, 0,
	)
	if err != nil {
		fmt.Println("Error inserting task:", err)
		return
	}

	fmt.Println("Task added:", taskName)
}
