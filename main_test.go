package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "basic_trimming_and_splitting",
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "longer_string_with_mixed_case",
			input:    "    this IS a longer! STRing!",
			expected: []string{"this", "is", "a", "longer!", "string!"},
		},
		{
			name:     "string_with_symbols",
			input:    "this string> Has > symbols",
			expected: []string{"this", "string>", "has", ">", "symbols"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := cleanInput(c.input)
			if len(actual) != len(c.expected) {
				t.Errorf("actual slice length: %d (expected: %d)", len(actual), len(c.expected))
				return
			}
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				if word != expectedWord {
					t.Errorf("word mismatch at index %d (found %s, expected %s)", i, word, expectedWord)
					return
				}
			}
		})
	}
}
