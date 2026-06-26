package commands

import (
	"errors"
	"fmt"
)

func exploreCommand(args []string) error {
	if len(args) == 0 {
		return errors.New("Must supply an area to explore")
	}
	area_name := args[0]
	fmt.Printf("Exploring %s...\n", area_name)
	url := LOCATION_AREA_ENDPOINT + "/" + area_name
	resp, err := GetResource[LocationArea](url)
	if err != nil {
		return err
	}
	if len(resp.Encounters) > 0 {
		fmt.Println("Found Pokemon: ")
		for _, encounter := range resp.Encounters {
			fmt.Printf(" - %s\n", encounter.Pokemon.Name)
		}
	}
	return nil
}
