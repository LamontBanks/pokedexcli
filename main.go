package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Commands
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Commands ---
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	// Print all the commands and descriptions
	var msg string
	for _, cmd := range commands {
		msg += fmt.Sprintf("%v: %v\n", cmd.name, cmd.description)
	}
	fmt.Println(msg)

	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

// Response Structs
// Converted using https://mholt.github.io/json-to-go/
// TODO: Move into internal package

// Endpoint: https://pokeapi.co/api/v2/location  (no "s" at end)
// Doc: https://pokeapi.co/docs/v2#locations
type Maps struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// GET Pokemon maps/locations
func mapCommand() error {
	// Create the URL
	fullUrl := "https://pokeapi.co/api/v2/location"

	// Create the request
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return err
	}

	// Create the client
	client := &http.Client{}

	// Make the request
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	// Ensure request is closed
	defer res.Body.Close()

	// Parse the location JSON
	var maps Maps
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&maps); err != nil {
		return err
	}

	// Print map names
	for _, location := range maps.Results {
		fmt.Println(location.Name)
	}
	return nil
}

// Parse functions ---
// Splits user input by whitespace,
// trims leading and trailing whitespace from tokens,
// normalizes token to lowercase
func cleanInput(text string) []string {
	var tokens []string

	splitTokens := strings.Fields(text)

	for _, token := range splitTokens {
		trimmedToken := strings.TrimRight(strings.TrimLeft(token, " "), " ")
		lowerCaseToken := strings.ToLower(trimmedToken)
		tokens = append(tokens, lowerCaseToken)
	}

	return tokens
}

// Main ---

var commands = map[string]cliCommand{}

func main() {
	// Set commands
	commands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	commands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	commands["map"] = cliCommand{
		name:        "map",
		description: "List Pokemon locations",
		callback:    mapCommand,
	}

	// Read-Eval-Print-Loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")

		scanner.Scan()
		userInput := scanner.Text()
		tokens := cleanInput(userInput)

		// Run command
		cmd, exists := commands[tokens[0]]
		if !exists {
			fmt.Println("Unknown command:", tokens[0])
			continue
		}
		if err := cmd.callback(); err != nil {
			fmt.Println(err)
		}
	}
}
