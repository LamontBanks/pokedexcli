package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")

		scanner.Scan()

		userInput := scanner.Text()
		tokens := cleanInput(userInput)

		fmt.Printf("Your command was: %v\n", tokens[0])
	}
}

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
