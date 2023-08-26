package api

import (
	"encoding/json"
	"io"
)

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationAreasResponse, error) {
	if body, inCache := client.cache.Get(url); inCache {
		var locAreas LocationAreasResponse
		json.Unmarshal(body, &locAreas)

		return locAreas, nil
	}

	res, err := client.http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	go client.cache.Add(url, body)

	var locAreas LocationAreasResponse
	json.Unmarshal(body, &locAreas)

	return locAreas, nil
}
