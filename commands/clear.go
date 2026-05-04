package commands

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "modernc.org/sqlite"
)

func Clear(db *sql.DB, args []string) {
	if len(args) > 1 {
		fmt.Println("Too many arguments.")
		return
	}

	fmt.Print("Are you sure you want to remove all tasks? (y/n): ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(strings.ToLower(input))

	if input == "y" || input == "yes" {
		_, err := db.Exec("DROP TABLE IF EXISTS tasks")
		if err != nil {
			fmt.Println("Query error:", err)
			return
		} else {
			fmt.Println("Succesfully removed all tasks")
			return
		}

	} else {
		fmt.Println("Operation cancelled.")
		return
	}
}
