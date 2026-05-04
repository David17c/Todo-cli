// main.go
package main

import (
	"os"
	"todo-cli/cli"
	"todo-cli/interactive"
)

func main() {
	if len(os.Args) < 2 {
		interactive.ReplMode()
		return
	}

	cli.Parser(os.Args[1:])
}
