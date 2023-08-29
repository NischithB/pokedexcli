package cli

import (
	"fmt"
	"log"
	"math/rand"
	"os"
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

	locations, err := cfg.services.GetLocationAreas(*cfg.nextLocationAreas)
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
	locations, err := cfg.services.GetLocationAreas(*cfg.prevLocationAreas)
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
	if len(args) != 1 {
		log.Printf("1 argument needed: 'explore {area}', %d were given", len(args))
		return
	}

	fmt.Printf("Exploring %s...\n", args[0])
	pokes, err := cfg.services.GetPokemonsInArea(args[0])

	if err != nil {
		log.Printf("failed to display pokemons in %s\n", args[0])
	}

	fmt.Println("Found Pokemon:")
	for _, poke := range pokes.PokemonEncounters {
		fmt.Printf("  -  %s\n", poke.Pokemon.Name)
	}

	return
}

func handleCatch(cfg *Config, args ...string) (err error) {
	if len(args) != 1 {
		log.Printf("1 argument needed: 'catch {pokemon}', %d were given", len(args))
		return
	}
	pokeName := args[0]

	poke, err := cfg.services.GetPokemon(pokeName)
	if err != nil {
		log.Println("failed to fetch pokemon")
		return
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokeName)
	xp := rand.Intn(poke.BaseExperience)

	if xp > 40 {
		fmt.Printf("%s escaped!\n", poke.Name)
		return nil
	}

	cfg.services.PokeStore[pokeName] = poke
	fmt.Printf("%s was caught!\n", poke.Name)
	return
}

func handleInspect(cfg *Config, args ...string) (err error) {
	if len(args) != 1 {
		log.Printf("1 argument needed: 'inspect {pokemon}', %d were given", len(args))
		return
	}

	pokeName := args[0]
	poke, isCaught := cfg.services.PokeStore[pokeName]
	if !isCaught {
		fmt.Println("you haven't caught this pokemon")
		return
	}

	// Display details
	fmt.Println("Name:", poke.Name)
	fmt.Println("Height:", poke.Height)
	fmt.Println("Weight:", poke.Weight)
	fmt.Println("Stats:")
	for _, stat := range poke.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range poke.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}

	return
}

func handleExit(cfg *Config, args ...string) (err error) {
	os.Exit(0)
	return
}
