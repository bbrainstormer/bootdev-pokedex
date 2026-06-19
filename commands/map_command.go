package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const LOCATION_AREA_ENDPOINT string = "https://pokeapi.co/api/v2/location-area"

type locationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type locationResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []locationArea `json:"results"`
}

func mapCommand() error {
	endpoint := globalConfig.next
	if endpoint == "" {
		endpoint = LOCATION_AREA_ENDPOINT
	}

	response, err := getLocationAreas(endpoint)
	if err != nil {
		return err
	}

	globalConfig.next = response.Next
	globalConfig.previous = response.Previous

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func mapBackCommand() error {
	endpoint := globalConfig.previous
	if endpoint == "" {
		endpoint = LOCATION_AREA_ENDPOINT
	}

	response, err := getLocationAreas(endpoint)
	if err != nil {
		return err
	}

	globalConfig.next = response.Next
	globalConfig.previous = response.Previous

	for _, result := range response.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func getLocationAreas(url string) (locationResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return locationResponse{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	decoded := locationResponse{}
	err = decoder.Decode(&decoded)
	if err != nil {
		return locationResponse{}, err
	}

	return decoded, nil
}
