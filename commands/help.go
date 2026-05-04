// commands/help.go
package commands

import (
	"fmt"
)

func Help() {
	fmt.Print("Usage: todo <command> [arguments]")
	fmt.Println("")
	fmt.Println("Commands:")
	fmt.Println("	add <task>        Add a new task")
	fmt.Println("	remove <id>       Remove a task by ID")
	fmt.Println("	list [done|todo]  List tasks (optionally filter by status)")
	fmt.Println("	edit <id>         Edit an existing task")
	fmt.Println("	mark <id>         Mark a task as done")
	fmt.Println("	unmark <id>       Mark a task as not done")
	fmt.Println("	help              Show this help message")
}
