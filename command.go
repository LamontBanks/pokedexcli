package main

import "github.com/LamontBanks/pokedexcli/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}
