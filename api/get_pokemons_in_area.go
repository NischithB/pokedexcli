package api

import (
	"encoding/json"
	"fmt"
	"io"
)

type PokemonsInAreaResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetPokemonsInArea(area string) (PokemonsInAreaResponse, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", area)

	if body, inCache := client.cache.Get(url); inCache {
		var pokes PokemonsInAreaResponse
		json.Unmarshal(body, &pokes)

		return pokes, nil
	}

	res, err := client.http.Get(url)
	if err != nil {
		return PokemonsInAreaResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return PokemonsInAreaResponse{}, err
	}
	go client.cache.Add(url, body)

	var pokes PokemonsInAreaResponse
	json.Unmarshal(body, &pokes)
	return pokes, nil
}
