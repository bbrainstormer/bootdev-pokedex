package commands

import (
	"fmt"
)

const LOCATION_AREA_ENDPOINT string = "https://pokeapi.co/api/v2/location-area"

func mapCommand(args []string) error {
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

func mapBackCommand(args []string) error {
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
	return GetResource[locationResponse](url)
}
