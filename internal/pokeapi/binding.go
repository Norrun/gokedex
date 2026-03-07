package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AreaLocationResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Area `json:"results"`
}
type Area struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetAreas(altUrl string) (AreaLocationResult, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if altUrl != "" {
		url = altUrl
	}
	body, err := callAPI(url)
	if err != nil {
		return AreaLocationResult{}, err
	}
	areaLocation := AreaLocationResult{}
	err = json.Unmarshal(body, &areaLocation)
	if err != nil {
		return AreaLocationResult{}, err
	}
	return areaLocation, nil
}

func callAPI(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}
