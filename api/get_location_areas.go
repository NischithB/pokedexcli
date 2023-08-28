package api

import (
	"encoding/json"
	"io"
)

func (services *Services) GetLocationAreas(url string) (LocationAreasResponse, error) {
	if body, inCache := services.Cache.Get(url); inCache {
		var locAreas LocationAreasResponse
		json.Unmarshal(body, &locAreas)

		return locAreas, nil
	}

	res, err := services.HttpClient.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	go services.Cache.Add(url, body)

	var locAreas LocationAreasResponse
	json.Unmarshal(body, &locAreas)

	return locAreas, nil
}

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
