// main.go
package main

import (
	"fmt"
	"os"
	"todo-cli/commands"

	_ "modernc.org/sqlite"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "please provide a command (add, rem, list, edit)")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		commands.Add(os.Args[2:])
	case "remove":
		commands.Remove(os.Args[2:])
	case "list":
		commands.List(os.Args[2:])
	case "edit":
		commands.Edit(os.Args[2:])
	case "help":
		commands.Help()
	case "mark":
		commands.Mark(os.Args[2:])
	case "unmark":
		commands.Unmark(os.Args[2:])
	default:
		fmt.Fprintln(os.Stderr, "unknown command:", cmd)
		os.Exit(1)
	}
}
