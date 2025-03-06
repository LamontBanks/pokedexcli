package main

import httpconfig "github.com/LamontBanks/pokedexcli/internal/http_config"

type cliCommand struct {
	name        string
	description string
	callback    func(config *httpconfig.Config, args []string) error
}
