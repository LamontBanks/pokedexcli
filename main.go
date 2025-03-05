package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Commands
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// Command Functions ---
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
