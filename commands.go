package main

import (
	"fmt"
	"os"
)

/*var registry = */
func getRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func commandExit() error {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return err
}

func commandHelp() error {
	usage := ""
	for _, command := range getRegistry() {
		usage += fmt.Sprintf("%s: %s\n", command.name, command.description)

	}
	_, err := fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n%s", usage)
	return err
}
