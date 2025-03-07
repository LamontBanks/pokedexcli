package showmaps

import (
	"fmt"

	"github.com/LamontBanks/pokedexcli/internal/pokeapi"
)

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

// GET Pokemon maps or move FORWARD through pages of results
func MapCommand(config *pokeapi.Config, args []string) error {
	fullUrl := "https://pokeapi.co/api/v2/location-area"

	// Use the pagination url over the default, if set
	if config.NextUrl != nil {
		fullUrl = *config.NextUrl
	}

	// Make request
	var mapsResponse Maps
	if err := pokeapi.PokeCacheHttpGet(fullUrl, &mapsResponse, config); err != nil {
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
func MapBackCommand(config *pokeapi.Config, args []string) error {
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
	if err := pokeapi.PokeCacheHttpGet(fullUrl, &mapsResponse, config); err != nil {
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
