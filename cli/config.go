package cli

import (
	"github.com/NischithB/pokedexcli/api"
)

type Config struct {
	services          api.Services
	nextLocationAreas *string
	prevLocationAreas *string
}
