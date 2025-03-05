package main

import (
	"fmt"
	"os"

	httpconfig "github.com/LamontBanks/pokedexcli/internal/http_config"
)

func commandExit(config *httpconfig.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
