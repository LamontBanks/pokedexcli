package httpconfig

import (
	"github.com/LamontBanks/pokedexcli/internal/pokecache"
	"github.com/LamontBanks/pokedexcli/internal/pokedex"
)

// Struct to save info between HTTP calls, ex: pagination variables
type Config struct {
	NextUrl     *string
	PreviousUrl *string
	Cache       pokecache.Cache
	Pokedex     pokedex.Pokedex
}
