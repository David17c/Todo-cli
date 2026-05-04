// commands/help.go
package commands

import (
	"fmt"
)

func Help() {
	fmt.Println(`Usage: todo <command> [arguments]

Commands:
  add <task>        	Add a new task
  remove <id>       	Remove a task by ID
  list [done|todo]  	List tasks
  edit <id> <new task>  Edit an existing task
  mark <id>         	Mark a task as done
  unmark <id>       	Mark a task as not done
  help              	Show this help message`)
}
