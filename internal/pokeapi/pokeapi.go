package pokeapi

import (
	"fmt"

	httpclient "github.com/LamontBanks/pokedexcli/internal/http_client"
	httpconfig "github.com/LamontBanks/pokedexcli/internal/http_config"
)

// Response Structs
// Converted using https://mholt.github.io/json-to-go/
// Manually set nullable strings to `*string``

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

// GET Pokemon maps, move FORWARD through pages of results
func MapCommand(config *httpconfig.Config) error {
	// Go to the base URL, or next page (if set)
	fullUrl := "https://pokeapi.co/api/v2/location-area"
	if config.NextUrl != nil {
		fullUrl = *config.NextUrl
	}

	var maps Maps
	httpclient.Get(fullUrl, &maps)

	// Save the updated Previous and Next URLs of pagination
	config.PreviousUrl = maps.Previous
	config.NextUrl = maps.Next

	// Print map names
	for _, location := range maps.Results {
		fmt.Println(location.Name)
	}
	return nil
}

// GET Pokemon maps, move BACKWARDS through results
// TODO - make into singular function
func MapBackCommand(config *httpconfig.Config) error {
	// Go to the base URL, or next page (if set)
	fullUrl := ""
	if config.PreviousUrl != nil {
		fullUrl = *config.PreviousUrl
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	// Populate the response in `maps`
	var maps Maps
	httpclient.Get(fullUrl, &maps)

	// Save the updated Previous, Next URLs of pagination
	config.PreviousUrl = maps.Previous
	config.NextUrl = maps.Next

	// Print map names
	for _, location := range maps.Results {
		fmt.Println(location.Name)
	}
	return nil
}
