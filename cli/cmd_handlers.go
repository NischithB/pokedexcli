package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/NischithB/pokedexcli/api"
)

func handleHelp(cfg *Config, args ...string) (err error) {
	fmt.Print("\nWelcome to Pokedex\n")
	fmt.Print("\nUsage:\n\n")

	for key, val := range getCommands() {
		fmt.Printf("%s: %s\n", key, val.description)
	}

	fmt.Println()
	return
}

func handleMap(cfg *Config, args ...string) (err error) {
	if cfg.nextLocationAreas == nil {
		fmt.Println("You have reached the end")
		return
	}

	locations, err := api.GetLocationAreas(*cfg.nextLocationAreas)
	if err != nil {
		log.Printf("failed to display location areas: %v", err)
		return
	}

	// Update config
	cfg.nextLocationAreas = locations.Next
	cfg.prevLocationAreas = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()

	return
}

func handleMapb(cfg *Config, args ...string) (err error) {
	if cfg.prevLocationAreas == nil {
		fmt.Println("You are at start")
		return
	}
	locations, err := api.GetLocationAreas(*cfg.prevLocationAreas)
	if err != nil {
		log.Printf("failed to display location areas: %v", err)
		return
	}

	// Update config
	cfg.nextLocationAreas = locations.Next
	cfg.prevLocationAreas = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	fmt.Println()

	return
}

func handleExplore(cfg *Config, args ...string) (err error) {
	if len(args) < 1 {
		log.Printf("'area' is missing, 1 argument needed: 'explore {area}'")
		return
	}

	pokes, err := api.GetPokemonsInArea(args[0])

	if err != nil {
		log.Printf("failed to display pokemons in %s", args[0])
	}

	for _, poke := range pokes.PokemonEncounters {
		fmt.Println(poke.Pokemon.Name)
	}

	return
}

func handleExit(cfg *Config, args ...string) (err error) {
	os.Exit(0)
	return
}
