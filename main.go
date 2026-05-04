// main.go
package main

import (
	"fmt"
	"os"

	"todo-cli/commands"
	"todo-cli/db"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "please provide a command (add, rem, list, edit)")
		os.Exit(1)
	}

	cmd := os.Args[1]

	dbConn, err := db.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to initialize database:", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	switch cmd {
	case "add":
		commands.Add(dbConn, os.Args[2:])
	case "remove":
		commands.Remove(dbConn, os.Args[2:])
	case "list":
		commands.List(dbConn, os.Args[2:])
	case "edit":
		commands.Edit(dbConn, os.Args)
	case "help":
		commands.Help()
	case "mark":
		commands.Mark(dbConn, os.Args[2:])
	case "unmark":
		commands.Unmark(dbConn, os.Args[2:])
	default:
		fmt.Fprintln(os.Stderr, "unknown command:", cmd)
		os.Exit(1)
	}
}
