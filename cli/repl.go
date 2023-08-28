package cli

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/NischithB/pokedexcli/api"
	"github.com/NischithB/pokedexcli/cache"
	"github.com/NischithB/pokedexcli/utils"
)

func StartREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	url := "https://pokeapi.co/api/v2/location-area"
	config := Config{
		services: api.Services{
			HttpClient: http.Client{Timeout: time.Second * time.Duration(5)},
			Cache:      cache.NewCache(time.Second * time.Duration(20)),
			PokeStore:  map[string]api.Pokemon{},
		},
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

		args := cmdLine[1:]
		if err := cmd.callback(&config, args...); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
