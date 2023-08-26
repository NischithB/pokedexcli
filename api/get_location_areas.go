package api

import (
	"encoding/json"
	"io"
	"net/http"
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
	res, err := http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	var locAreas LocationAreasResponse
	json.Unmarshal(body, &locAreas)

	return locAreas, nil
}
