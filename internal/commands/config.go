package commands

import (
	"time"

	"github.com/bbrainstormer/bootdev-pokedex/internal/lib"
	"github.com/bbrainstormer/bootdev-pokedex/internal/pokecache"
)

type Config struct {
	next          string
	previous      string
	Cache         *pokecache.Cache
	CaughtPokemon lib.MutexMap[string, Pokemon]
}

type CaughtPokemon struct {
}

var globalConfig Config = Config{
	Cache:         pokecache.NewCache(time.Minute),
	CaughtPokemon: lib.NewMutexMap[string, Pokemon](),
}
