package main

import (
	"fmt"

	httpconfig "github.com/LamontBanks/pokedexcli/internal/http_config"
)

func commandHelp(config *httpconfig.Config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	// Print all the commands and descriptions
	var msg string
	for _, cmd := range getCommands() {
		msg += fmt.Sprintf("%v: %v\n", cmd.name, cmd.description)
	}
	fmt.Println(msg)

	return nil
}
