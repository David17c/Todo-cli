// commands/edit.go
package commands

import (
	//	"database/sql"
	"database/sql"
	"fmt"
	"strings"
	// _ "modernc.org/sqlite"
)

func Edit(db *sql.DB, args []string) {
	task := strings.Join(args, " ")
	fmt.Println(task)
}

//Coming soon
