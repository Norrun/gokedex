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
