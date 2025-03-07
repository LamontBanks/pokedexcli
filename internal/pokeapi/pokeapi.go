// TODO - Move functions into own file
package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"time"

	httpclient "github.com/LamontBanks/pokedexcli/internal/http-client"
	"github.com/LamontBanks/pokedexcli/internal/pokecache"
)

// Saves select data from PokeAPI calls
type Config struct {
	NextUrl     *string
	PreviousUrl *string
	Cache       pokecache.Cache
	Pokedex     map[string]Pokemon
}

// Response Structs
// Converted using https://mholt.github.io/json-to-go/
// Manually set nullable strings to `*string``

// Location-area without parameters
// https://pokeapi.co/api/v2/location-area
// https://pokeapi.co/docs/v2#location-areas
type Maps struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

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

// https://pokeapi.co/docs/v2#pokemon
// Minimum fields needed
type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Order          int    `json:"order"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

// GET Pokemon maps or move FORWARD through pages of results
func MapCommand(config *Config, args []string) error {
	fullUrl := "https://pokeapi.co/api/v2/location-area"

	// Use the pagination url over the default, if set
	if config.NextUrl != nil {
		fullUrl = *config.NextUrl
	}

	// Make request
	var mapsResponse Maps
	if err := PokeCacheHttpGet(fullUrl, &mapsResponse, config); err != nil {
		return err
	}

	// Save the updated Previous and Next URLs of pagination
	config.PreviousUrl = mapsResponse.Previous
	config.NextUrl = mapsResponse.Next

	// Print map names
	for _, location := range mapsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

// GET Pokemon maps, move BACKWARDS through results
func MapBackCommand(config *Config, args []string) error {
	// Go to the base URL, or next page (if set)
	fullUrl := ""
	if config.PreviousUrl != nil {
		fullUrl = *config.PreviousUrl
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	// Make request, cache
	var mapsResponse Maps
	if err := PokeCacheHttpGet(fullUrl, &mapsResponse, config); err != nil {
		return err
	}

	// Save the updated Previous, Next URLs of pagination
	config.PreviousUrl = mapsResponse.Previous
	config.NextUrl = mapsResponse.Next

	// Print map names
	for _, location := range mapsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

// List Pokemon found the given emap
func ExploreMapCommand(config *Config, args []string) error {
	fullUrl := "https://pokeapi.co/api/v2/location-area"

	// The second token should be the name or id parameter
	if len(args) < 2 {
		return errors.New("missing id/name parameter")
	}
	fullUrl += fmt.Sprintf("/%v", args[1])

	// Make request, cache
	var locationAreaResponse LocationArea
	if err := PokeCacheHttpGet(fullUrl, &locationAreaResponse, config); err != nil {
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

// Attempt to catch the Pokemon
// Add to Pokedex if caught
func CatchCommand(config *Config, args []string) error {
	fullUrl := "https://pokeapi.co/api/v2/pokemon"

	// The second token should be the name or id parameter
	if len(args) < 2 {
		return errors.New("missing pokemon name name/id")
	}
	fullUrl += fmt.Sprintf("/%v", args[1])

	// Make request
	var pokemonResponse Pokemon
	if err := PokeCacheHttpGet(fullUrl, &pokemonResponse, config); err != nil {
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

// HTTP GET saving/pulling from the cache
// Unmarshal's the JSON response into provided response
func PokeCacheHttpGet(url string, response any, config *Config) error {
	// Check cache first
	cachedBytes, responseIsCached := config.Cache.Get(url)

	if responseIsCached {
		// Unmarshal the cached bytes into the response struct
		if err := json.Unmarshal(cachedBytes, &response); err != nil {
			return err
		}
	} else {
		// Otherwise, make the actual request
		if err := httpclient.Get(url, &response); err != nil {
			return err
		}

		// Convert to a []byte
		encodededBytes, err := json.Marshal(response)
		if err != nil {
			return err
		}

		// Then save to the cache
		config.Cache.Add(url, encodededBytes)
	}

	return nil
}
