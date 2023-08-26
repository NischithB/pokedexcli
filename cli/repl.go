package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/NischithB/pokedexcli/utils"
)

func StartREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	url := "https://pokeapi.co/api/v2/location-area"
	config := Config{
		nextLocationAreas: &url,
		prevLocationAreas: nil,
	}

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()

		cmdLine := utils.FormatCmd(scanner.Text())

		if len(cmdLine) == 0 {
			continue
		}

		cmd, ok := commands[cmdLine[0]]
		if !ok {
			fmt.Println("invalid command: try 'help' to know more about pokedex")
			continue
		}

		if err := cmd.callback(&config); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
