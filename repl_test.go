package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "greet  your    self",
			expected: []string{"greet", "your", "self"},
		},
		{
			input:    "Not-Adding-Whitespace",
			expected: []string{"not-adding-whitespace"},
		},
		// add more cases here
	}
	// loop over the cases and run the tests:
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Fatalf("Length mismatch between expected and actual return values")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected %v but got %v", expectedWord, word)
			}
		}
	}
}
