// commands/edit.go
package commands

import (
	//	"database/sql"
	"fmt"
	"strings"
	// _ "modernc.org/sqlite"
)

func Edit(args []string) {
	task := strings.Join(args, " ")
	fmt.Println(task)
}
