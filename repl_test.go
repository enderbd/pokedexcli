package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{		
			input: " what    have you EATEN",
			expected: []string{"what", "have", "you", "eaten"},
		},
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Actual lenghts is different from expected:\n actual len %v != expected len %v", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Actual word is not the expected one: %v != %v", word, expectedWord)
			}
		}
	}
}
