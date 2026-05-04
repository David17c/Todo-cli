// commands/list.go
package commands

import (
	"database/sql"
	"fmt"
	"log"
)

func List(args []string) {
	db, err := sql.Open("sqlite", "file:tasks.db")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	query := "SELECT id, taskname, status FROM tasks"
	var queryArgs []any

	if len(args) > 1 {
		fmt.Println("too many arguments.")
		return
	}

	if len(args) == 1 {
		switch args[0] {
		case "done":
			query += " WHERE status = ?"
			queryArgs = append(queryArgs, 1)
		case "todo":
			query += " WHERE status = ?"
			queryArgs = append(queryArgs, 0)
		default:
			fmt.Println("invalid filter argument.")
			return
		}
	}

	rows, err := db.Query(query, queryArgs...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var i = 1
	for rows.Next() {
		var id int
		var taskname string
		var status int

		if err := rows.Scan(&id, &taskname, &status); err != nil {
			log.Fatal(err)
		}

		check := " "
		if status == 1 {
			check = "x"
		}

		fmt.Printf("%-3d: [%-1s] %-20s (id:%3d)\n", i, check, taskname, id)
		i++
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
