package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name		string
		input 		string
		expected	[]string
	}{
		{
			name:		"empty string",
			input: 		"",
			expected: 	[]string{},
		},
		{
			name: 		"all small letters",
			input:		"hello world",
			expected:	[]string{"hello", "world"},
		},
		{
			name: 		"all capital letters",
			input:		"HELLO WORLD",
			expected:	[]string{"hello", "world"},
		},
		{
			name: 		"camel cs",
			input:		"hELLO wORLD",
			expected:	[]string{"hello", "world"},
		},
		{
			name: 		"title cs",
			input:		"Hello World",
			expected:	[]string{"hello", "world"},
		},
		{
			name: 		"symbols included",
			input:		"Hello World!",
			expected:	[]string{"hello", "world!"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf(
				"length of the output do not match: Actual '%v' vs Expected '%v'", len(actual), len(c.expected))
		}
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf(
				"result do not match: Actual '%v' vs Expected '%v'", actual[i], c.expected[i])
			}
		}
	}

}

