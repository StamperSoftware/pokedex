package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	input string
	expected []string
}

func TestCleanInput(t *testing.T) {
	
	pass := 0
	fails := 0

	cases := []testCase{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "HELLO WORLD",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		hasFailed := false
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			fails++
			t.Errorf("Failed: %d\n", len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				fails++
				hasFailed = true
				t.Errorf("Failed: %s\n", word)
				break
			}
		}
		if !hasFailed {
			pass++
		}
	}

	fmt.Printf("Success Count: %d\n", pass)
	fmt.Printf("Failed Count: %d\n", fails)

}


