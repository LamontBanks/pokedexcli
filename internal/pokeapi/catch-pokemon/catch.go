package catchpokemon

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/LamontBanks/pokedexcli/internal/pokeapi"
)

// Attempt to catch the Pokemon
// Add to Pokedex if caught
func CatchCommand(config *pokeapi.Config, args []string) error {
	fullUrl := "https://pokeapi.co/api/v2/pokemon"

	// The second token should be the name or id parameter
	if len(args) < 2 {
		return errors.New("missing pokemon name name/id")
	}
	fullUrl += fmt.Sprintf("/%v", args[1])

	// Make request
	var pokemonResponse pokeapi.Pokemon
	if err := pokeapi.PokeCacheHttpGet(fullUrl, &pokemonResponse, config); err != nil {
		return err
	}

	// Attempt to catch pokemon
	// Determine catch change from base_experience
	// Flat 30% chance
	fmt.Printf("Throwing a Pokeball at %v...\n", pokemonResponse.Name)

	catchChance := rand.Intn(100)
	const waitTime = 500 * time.Millisecond
	time.Sleep(waitTime)

	if catchChance >= 30 {
		fmt.Printf("Great! Caught %v\n", pokemonResponse.Name)
		time.Sleep(waitTime)

		fmt.Println("Registering to the Pokedex...")
		config.Pokedex[pokemonResponse.Name] = pokemonResponse
		time.Sleep(waitTime)
	} else {
		time.Sleep(waitTime)
		fmt.Println("Catch failed!")
	}

	return nil
}
