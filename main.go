package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/LamontBanks/pokedexcli/internal/pokeapi"
	inspectpokemon "github.com/LamontBanks/pokedexcli/internal/pokeapi/inspect-pokemon"
	"github.com/LamontBanks/pokedexcli/internal/pokeapi/pokedex"
	"github.com/LamontBanks/pokedexcli/internal/pokecache"
)

// Main ---
var commands = map[string]cliCommand{}

func main() {
	// CLI commands
	setCommands()

	// Config needed for API calls
	commandConfig := pokeapi.Config{
		PreviousUrl: nil,
		NextUrl:     nil,
		Cache:       pokecache.NewCache(5 * time.Minute),
		Pokedex:     map[string]pokeapi.Pokemon{},
	}

	// Read-Eval-Print-Loop
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
		// If valid, run the command
		if err := cmd.callback(&commandConfig, tokens); err != nil {
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

func setCommands() {
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
	commands["explore"] = cliCommand{
		name:        "explore <map>",
		description: "List Pokemon found in given location",
		callback:    pokeapi.ExploreMapCommand,
	}
	commands["catch"] = cliCommand{
		name:        "catch <pokemon>",
		description: "Attempts to catch the Pokemon, saves to the Pokedex",
		callback:    pokeapi.CatchCommand,
	}
	commands["inspect"] = cliCommand{
		name:        "inspect <pokemon>",
		description: "Show data for a caught Pokemon",
		callback:    inspectpokemon.InspectCommand,
	}
	commands["pokedex"] = cliCommand{
		name:        "pokedex",
		description: "Show data for all caught Pokemon",
		callback:    pokedex.PokedexCommand,
	}
}
