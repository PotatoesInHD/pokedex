package main

import "testing"


func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}	{
		{
		input: "  Hello  world  ",
		expected: []string{"hello", "world"},
		},
		{
		input: "  Hello  world d@dsads ",
		expected: []string{"hello", "world", "d@dsads"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("expected length: %v, got: %v", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				 t.Errorf("expected: %s, got: %s", expectedWord, word)
			}
		}
	}
}
