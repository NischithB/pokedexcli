package api

import (
	"net/http"
	"time"

	"github.com/NischithB/pokedexcli/cache"
)

type APIClient struct {
	http  http.Client
	cache cache.Cache
}

var client = APIClient{
	http:  http.Client{Timeout: time.Second * time.Duration(5)},
	cache: cache.NewCache(time.Second * time.Duration(20)),
}
