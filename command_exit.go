package main

import (
	"fmt"
	"os"

	"github.com/LamontBanks/pokedexcli/internal/pokeapi"
)

func commandExit(config *pokeapi.Config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
