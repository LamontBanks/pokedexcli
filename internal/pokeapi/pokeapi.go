// TODO - Move functions into own file
package pokeapi

import (
	"encoding/json"

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
