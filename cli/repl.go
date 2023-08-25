package cli

import (
	"bufio"
	"fmt"
	"os"
)

func StartREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		cmdLine := formatCmd(scanner.Text())

		if len(cmdLine) == 0 {
			continue
		}

		cmd, ok := commands[cmdLine[0]]
		if !ok {
			fmt.Println("invalid command: try 'help' to know more about pokedex")
			continue
		}

		if err := cmd.Callback(); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
