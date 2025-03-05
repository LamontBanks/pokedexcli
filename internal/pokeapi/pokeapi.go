package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Used to save some response information to use between requests
type Config struct {
	NextUrl     string
	PreviousUrl string
}

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
func MapCommand(config *Config) error {
	// Go to the base URL, or next page (if set)
	fullUrl := "https://pokeapi.co/api/v2/location-area"
	if config.NextUrl != "" {
		fullUrl = config.NextUrl
	}

	// Make the request
	resp, err := http.Get(fullUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Unmarshal the JSON response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var maps Maps
	if err := json.Unmarshal(data, &maps); err != nil {
		return err
	}

	// Save the updated Previous, Next URLs of pagination
	if maps.Previous != nil {
		config.PreviousUrl = *maps.Previous
	}
	if maps.Next != nil {
		config.NextUrl = *maps.Next
	}

	// Print map names
	for _, location := range maps.Results {
		fmt.Println(location.Name)
	}
	return nil
}

// GET Pokemon maps, move BACKWARDS through results
// TODO - make into singular function
func MapBackCommand(config *Config) error {
	// Go to the base URL, or next page (if set)
	fullUrl := ""
	if config.PreviousUrl != "" {
		fullUrl = config.PreviousUrl
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	// Make the request
	resp, err := http.Get(fullUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Unmarshal the JSON response body
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var maps Maps
	if err := json.Unmarshal(data, &maps); err != nil {
		return err
	}

	// Save the updated Previous, Next URLs of pagination
	if maps.Previous != nil {
		config.PreviousUrl = *maps.Previous
	} else {
		config.PreviousUrl = ""
	}

	if maps.Next != nil {
		config.NextUrl = *maps.Next
	} else {
		config.NextUrl = ""
	}

	// Print map names
	for _, location := range maps.Results {
		fmt.Println(location.Name)
	}
	return nil
}
