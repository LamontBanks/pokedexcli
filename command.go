package main

import httpconfig "github.com/LamontBanks/pokedexcli/internal/http_config"

type cliCommand struct {
	name        string
	description string
	callback    func(*httpconfig.Config) error
}
