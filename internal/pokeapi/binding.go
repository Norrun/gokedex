package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Norrun/gokedex/internal/pokecache"
)

var cache pokecache.Cache = pokecache.NewCache(time.Second * 5)

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
func GetArea(altBaseUrl, area string) (AreaLocation, error) {
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	if altBaseUrl != "" {
		baseUrl = altBaseUrl
	}
	url := fmt.Sprintf("%s%s", baseUrl, area)
	data, err := callAPI(url)
	if err != nil {
		return AreaLocation{}, err
	}
	areaLocation := AreaLocation{}
	err = json.Unmarshal(data, &areaLocation)
	if err != nil {
		return AreaLocation{}, err
	}
	return areaLocation, nil
}

func callAPI(url string) ([]byte, error) {
	body, exists := cache.Get(url)
	if exists {
		return body, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	body, err = io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("API call failed with: \n status code: %d \nbody: %s\n Make sure you use the correct command and arguments.\n", res.StatusCode, body)
	}
	if err != nil {
		return nil, err
	}
	cache.Add(url, body)
	return body, nil
}
