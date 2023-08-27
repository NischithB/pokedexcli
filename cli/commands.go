package cli

type Command struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

type Config struct {
	nextLocationAreas *string
	prevLocationAreas *string
}

func getCommands() map[string]Command {
	return map[string]Command{
		"help": {
			name:        "help",
			description: "Displays help message",
			callback:    handleHelp,
		},
		"map": {
			name:        "map",
			description: "Displays next set of location areas",
			callback:    handleMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous set of location areas",
			callback:    handleMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays all the Pokemon in the given area",
			callback:    handleExplore,
		},
		"exit": {
			name:        "exit",
			description: "Exits Pokedex",
			callback:    handleExit,
		},
	}
}
