package inspectpokemon

import (
	"errors"
	"fmt"

	"github.com/LamontBanks/pokedexcli/internal/pokeapi"
)

func InspectCommand(config *pokeapi.Config, args []string) error {
	// The second token should be the name or id parameter
	if len(args) < 2 {
		return errors.New("missing Pokemon's name")
	}

	// Ask Dexter for data
	pokemon, exists := config.Pokedex[args[1]]
	if !exists {
		return nil
	}

	// Format the output
	var formattedData string
	formattedData += fmt.Sprintf("Name: %v\n", pokemon.Name)
	formattedData += fmt.Sprintf("Height: %v\n", pokemon.Height)
	formattedData += fmt.Sprintf("Weight: %v\n", pokemon.Weight)

	formattedData += fmt.Sprintln("Stats:")
	formattedData += fmt.Sprintf("  -hp: %v\n", getPokemonStat("hp", pokemon))
	formattedData += fmt.Sprintf("  -attack: %v\n", getPokemonStat("attack", pokemon))
	formattedData += fmt.Sprintf("  -defense: %v\n", getPokemonStat("defense", pokemon))
	formattedData += fmt.Sprintf("  -special-attack: %v\n", getPokemonStat("special-attack", pokemon))
	formattedData += fmt.Sprintf("  -special-defense: %v\n", getPokemonStat("special-defense", pokemon))
	formattedData += fmt.Sprintf("  -speed: %v\n", getPokemonStat("speed", pokemon))

	formattedData += fmt.Sprintln("Types:")
	for _, pokemonType := range pokemon.Types {
		formattedData += fmt.Sprintf("  - %v\n", pokemonType.Type.Name)
	}

	fmt.Println(formattedData)

	return nil
}

func getPokemonStat(statName string, pokemon pokeapi.Pokemon) int {
	for _, stat := range pokemon.Stats {
		if stat.Stat.Name == statName {
			return stat.BaseStat
		}
	}
	return 0
}
