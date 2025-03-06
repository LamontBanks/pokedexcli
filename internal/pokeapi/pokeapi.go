package pokeapi

import (
	"encoding/json"
	"fmt"
	"slices"

	httpclient "github.com/LamontBanks/pokedexcli/internal/http_client"
	httpconfig "github.com/LamontBanks/pokedexcli/internal/http_config"
)

// Response Structs
// Converted using https://mholt.github.io/json-to-go/
// Manually set nullable strings to `*string``

// Location-area without parameters
// Endpoint: https://pokeapi.co/api/v2/location-area
// Doc: https://pokeapi.co/docs/v2#location-areas
type Maps struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// Endpoint: https://pokeapi.co/api/v2/location-area
// Doc: https://pokeapi.co/docs/v2#location-areas
// Keeping only the JSON fields we care about
type LocationArea struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// GET Pokemon maps or move FORWARD through pages of results
func MapCommand(config *httpconfig.Config) error {
	fullUrl := "https://pokeapi.co/api/v2/location-area"

	// Use the pagination url over the default, if set
	if config.NextUrl != nil {
		fullUrl = *config.NextUrl
	}

	// Check cache first, initialize struct to capture response
	cachedBytes, responseIsCached := config.Cache.Get(fullUrl)

	var mapsResponse Maps
	if responseIsCached {
		// Unmarshal the cached bytes into the response struct
		if err := json.Unmarshal(cachedBytes, &mapsResponse); err != nil {
			return err
		}
	} else {
		// Otherwise, make the actual request
		httpclient.Get(fullUrl, &mapsResponse)

		// Convert to a []byte
		encodededBytes, err := json.Marshal(mapsResponse)
		if err != nil {
			return err
		}

		// Then save to the cache
		config.Cache.Add(fullUrl, encodededBytes)
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
func MapBackCommand(config *httpconfig.Config) error {
	// Go to the base URL, or next page (if set)
	fullUrl := ""
	if config.PreviousUrl != nil {
		fullUrl = *config.PreviousUrl
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	// Check cache first
	cachedBytes, responseIsCached := config.Cache.Get(fullUrl)
	var mapsResponse Maps

	if responseIsCached {
		// Unmarshal the cached bytes into the response struct
		if err := json.Unmarshal(cachedBytes, &mapsResponse); err != nil {
			return err
		}
	} else {
		// Otherwise, make the actual request, then save to the cache
		httpclient.Get(fullUrl, &mapsResponse)

		encodededBytes, err := json.Marshal(mapsResponse)
		if err != nil {
			return err
		}
		config.Cache.Add(fullUrl, encodededBytes)
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

func ExploreMapCommand(config *httpconfig.Config) error {
	fullUrl := "https://pokeapi.co/api/v2/location-area/johto-route-32-area"

	// Check cache first, initialize struct to capture response
	cachedBytes, responseIsCached := config.Cache.Get(fullUrl)

	var locationAreaResponse LocationArea
	if responseIsCached {
		// Unmarshal the cached bytes into the response struct
		if err := json.Unmarshal(cachedBytes, &locationAreaResponse); err != nil {
			return err
		}
	} else {
		// Otherwise, make the actual request
		httpclient.Get(fullUrl, &locationAreaResponse)

		// Convert to a []byte
		encodededBytes, err := json.Marshal(locationAreaResponse)
		if err != nil {
			return err
		}

		// Then save to the cache
		config.Cache.Add(fullUrl, encodededBytes)
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
