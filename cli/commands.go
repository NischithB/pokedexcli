package cli

import (
	"fmt"
	"os"
)

type Command struct {
	Name        string
	Description string
	Callback    func() error
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			Name:        "help",
			Description: "Displays help message",
			Callback:    handleHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exits Pokedex",
			Callback:    handleExit,
		},
	}
}

func handleHelp() (err error) {
	fmt.Print("\nWelcome to Pokedex\n")
	fmt.Print("\nUsage:\n\n")

	for key, val := range getCommands() {
		fmt.Printf("%s: %s\n", key, val.Description)
	}

	fmt.Println()
	return
}

func handleExit() (err error) {
	os.Exit(0)
	return
}
