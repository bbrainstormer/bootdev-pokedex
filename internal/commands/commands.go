package commands

import (
	"fmt"
	"os"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(args []string) error
}

func exitCommand(args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func helpCommand(args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")

	for name, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", name, cmd.Description)
	}
	return nil
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "Exit",
			Description: "Exit the pokedex",
			Callback:    exitCommand,
		},
		"help": {
			Name:        "Help",
			Description: "Displays a help message",
			Callback:    helpCommand,
		},
		"map": {
			Name:        "Map",
			Description: "Displays a page of locations",
			Callback:    mapCommand,
		},
		"mapb": {
			Name:        "Mapb",
			Description: "Displays the previous page of locations",
			Callback:    mapBackCommand,
		},
		"explore": {
			Name:        "Explore",
			Description: "Lists the possible pokemon encounters in an area",
			Callback:    exploreCommand,
		},
		"catch": {
			Name:        "Catch",
			Description: "Attempts to catch the given pokemon",
			Callback:    captureCommand,
		},
		"inspect": {
			Name:        "Inspect",
			Description: "Inspects the given pokemon (must be caught)",
			Callback:    inspectCommand,
		},
		"pokedex": {
			Name:        "Pokedex",
			Description: "Lists the captured pokemon",
			Callback:    pokedexCommand,
		},
	}
}

func GetCommand(name string) (CliCommand, bool) {
	cmd, exists := getCommands()[name]
	return cmd, exists
}
