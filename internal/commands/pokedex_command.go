package commands

import "fmt"

func pokedexCommand(args []string) error {
	fmt.Println("Your Pokedex:")
	for name, _ := range globalConfig.CaughtPokemon.Iter {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
