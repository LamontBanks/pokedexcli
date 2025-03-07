package exploremap

import (
	"errors"
	"fmt"
	"slices"

	"github.com/LamontBanks/pokedexcli/internal/pokeapi"
)

// Location-area with parameters
// https://pokeapi.co/api/v2/location-area
// Keeping only the JSON fields we care about
type LocationArea struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// List Pokemon found the given emap
func ExploreMapCommand(config *pokeapi.Config, args []string) error {
	fullUrl := "https://pokeapi.co/api/v2/location-area"

	// The second token should be the name or id parameter
	if len(args) < 2 {
		return errors.New("missing id/name parameter")
	}
	fullUrl += fmt.Sprintf("/%v", args[1])

	// Make request, cache
	var locationAreaResponse LocationArea
	if err := pokeapi.PokeCacheHttpGet(fullUrl, &locationAreaResponse, config); err != nil {
		return err
	}

	// List Pokemon names
	locationPokemon := []string{}
	for _, pokemon := range locationAreaResponse.PokemonEncounters {
		locationPokemon = append(locationPokemon, pokemon.Pokemon.Name)
	}
	slices.Sort(locationPokemon)
	for _, pokemon := range locationPokemon {
		fmt.Println(pokemon)
	}

	return nil
}
