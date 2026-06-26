package commands

import (
	"errors"
	"fmt"
	"math/rand"
	"net/url"
)

func captureCommand(args []string) error {
	const POKEMON_ENDPOINT string = "https://pokeapi.co/api/v2/pokemon/"
	const POKEMON_SPECIES_ENDPOINT string = "https://pokeapi.co/api/v2/pokemon-species/"
	if len(args) == 0 {
		return errors.New("Must provide a pokemon to catch")
	}

	species_name := args[0]
	fullURL, err := url.JoinPath(POKEMON_ENDPOINT, species_name)
	if err != nil {
		return err
	}

	pokemon, err := GetResource[Pokemon](fullURL)
	if err != nil {
		fmt.Println("Cannot find pokemon.")
		return nil
	}

	pokemon_species, err := GetResource[PokemonSpecies](pokemon.Species.Url)
	if err != nil {
		return err
	}
	caught := float32(pokemon_species.CaptureRate)/255. > rand.Float32()

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)
	if caught {
		fmt.Printf("%s was caught\n", pokemon.Name)
		if !globalConfig.CaughtPokemon.Has(pokemon.Name) {
			globalConfig.CaughtPokemon.Set(pokemon.Name, pokemon)
		}
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}
