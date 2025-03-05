package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	httpconfig "github.com/LamontBanks/pokedexcli/internal/http_config"
	"github.com/LamontBanks/pokedexcli/internal/pokeapi"
	"github.com/LamontBanks/pokedexcli/internal/pokecache"
)

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
		description: "List Pokemon locations, page forward through results",
		callback:    pokeapi.MapCommand,
	}
	commands["mapb"] = cliCommand{
		name:        "mapb",
		description: "List Pokemon locations, page backwards through results",
		callback:    pokeapi.MapBackCommand,
	}

	// Read-Eval-Print-Loop
	commandConfig := httpconfig.Config{
		PreviousUrl: nil,
		NextUrl:     nil,
		Cache:       pokecache.NewCache(10 * time.Second),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")

		// Read
		scanner.Scan()
		userInput := scanner.Text()
		tokens := cleanInput(userInput)

		// Eval: Check if command is valid
		if len(tokens) == 0 {
			continue
		}
		cmd, exists := commands[tokens[0]]
		if !exists {
			fmt.Println("Unknown command:", tokens[0])
			continue
		}

		// Print
		// If valid, run the command, passing a pointer to config to save parts of the response
		// for subsequent calls...?
		if err := cmd.callback(&commandConfig); err != nil {
			fmt.Println(err)
		}
	}
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

func getCommands() map[string]cliCommand {
	return commands
}
