package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	// Create a struct to hold the test cases (input, output)
	cases := []struct {
		input    string   // Full CLI input string
		expected []string // Cleaned out words/tokens
	}{
		{
			// Happy path
			input:    "  Hello World  ",
			expected: []string{"hello", "world"},
		},
		{
			// Empty input
			input:    " ",
			expected: []string{},
		},
		{
			// Special characters
			input:    "ABC!#%%def JN*(&L:) ",
			expected: []string{"abc!#%%def", "jn*(&l:)"},
		},
	}

	// Loops through the list of "test case" structs, passing the input into the function
	// then checking against the expected output
	// Precisely checking each apsect of the result allows us to provide better failure messages
	// compared to just doing a single slice comparison
	for _, c := range cases {
		actual := cleanInput(c.input)

		// Check the lengths
		if len(actual) != len(c.expected) {
			t.Errorf("Words not correctly split. Actual %v | Expected %v", actual, c.expected)
			return
		}

		// Check each word
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("Word not properly cleaned. Actual '%v'| Expected '%v'", actual[i], c.expected[i])
				return
			}
		}

	}
}
