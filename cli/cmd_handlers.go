package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/NischithB/pokedexcli/api"
)

func handleHelp(_ *Config) (err error) {
	fmt.Print("\nWelcome to Pokedex\n")
	fmt.Print("\nUsage:\n\n")

	for key, val := range getCommands() {
		fmt.Printf("%s: %s\n", key, val.description)
	}

	fmt.Println()
	return
}

func handleMap(cfg *Config) (err error) {
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

func handleMapb(cfg *Config) (err error) {
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

func handleExit(_ *Config) (err error) {
	os.Exit(0)
	return
}
