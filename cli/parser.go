package cli

import (
	"fmt"
	"os"
	"todo-cli/commands"
	"todo-cli/db"
)

func Parser(args []string) {
	dbConn, err := db.Init()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to initialize database:", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "no command provided")
		return
	}

	switch args[0] {
	case "add":
		commands.Add(dbConn, args[1:])
	case "remove":
		commands.Remove(dbConn, args[1:])
	case "list":
		commands.List(dbConn, args[1:])
	case "edit":
		commands.Edit(dbConn, args)
	case "help":
		commands.Help()
	case "mark":
		commands.Mark(dbConn, args[1:])
	case "unmark":
		commands.Unmark(dbConn, args[1:])
	case "clear":
		commands.Clear(dbConn, args[1:])
	default:
		fmt.Fprintln(os.Stderr, "unknown command:", args[0])
	}
}
