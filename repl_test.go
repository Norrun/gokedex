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
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		t.Logf("erronious result: %q", actual)
		if len(actual) != len(c.expected) {
			t.Errorf("repl cleanInput err: expected slice length: %d, actual slice length: %d", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("repl cleanInput err: expected; %s actual; %s ", expectedWord, word)
			}
		}

	}

}
