package main

import "strings"

func cleanInput(text string) []string {
	parts := strings.Split(strings.ToLower(text), " ")

	words := make([]string, 0, len(parts))
	for _, v := range parts {
		if v != "" && v != " " {
			words = append(words, v)
		}
	}
	return words
}

type commandConfig struct {
	state []string
	args  []string
}

type cliCommand struct {
	name        string
	description string
	callback    func(commandConfig) ([]string, error)
}
