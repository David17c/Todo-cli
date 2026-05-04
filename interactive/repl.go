package interactive

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo-cli/cli"
)

func ReplMode() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter 'q' to quit.")
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		if input == "q" {
			return
		}

		cli.Parser(args)
	}
}
