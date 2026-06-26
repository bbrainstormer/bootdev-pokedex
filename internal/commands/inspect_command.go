package commands

import (
	"errors"
	"fmt"
)

func inspectCommand(args []string) error {
	if len(args) == 0 {
		return errors.New("Must have a pokemon to inspect")
	}

	pokemon, exists := globalConfig.CaughtPokemon.Get(args[0])
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %dcm\n", pokemon.Height*10)
	fmt.Printf("Wieght: %dg\n", pokemon.Weight*100)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf(" - %s\n", pokeType.Type.Name)
	}
	return nil
}
