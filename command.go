package main

import "github.com/LamontBanks/pokedexcli/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	callback    func(config *pokeapi.Config, args []string) error
}
