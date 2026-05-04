// commands/help.go
package commands

import (
	"fmt"
)

func Help() {
	fmt.Println("Commands:")
	fmt.Println("Add <task> (Adds a task)")
	fmt.Println("Remove <id> (Removes a task)")
	fmt.Println("List <done / todo> (Lists all tasks)")
	fmt.Println("Edit <id> (Edits an existing task)")
	fmt.Println("Help (Lists all possible commands)")
	fmt.Println("Mark <id> (changes task status to done)")
	fmt.Println("Unmark <id> (reverts task status to todo)")
}
