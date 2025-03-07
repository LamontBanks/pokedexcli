package pokedex

import (
	"fmt"

	"github.com/LamontBanks/pokedexcli/internal/pokeapi"
)

// Prints all the Pokemon caught (no order)
func PokedexCommand(config *pokeapi.Config, args []string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.Pokedex {
		fmt.Printf(" - %v\n", pokemon.Name)
	}

	return nil
}
