package api

import (
	"net/http"

	"github.com/NischithB/pokedexcli/cache"
)

type Services struct {
	HttpClient http.Client
	Cache      cache.Cache
	PokeStore  map[string]Pokemon
}
