package main

import (
	"fmt"
	"os"

	"github.com/Norrun/gokedex/internal/pokeapi"
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
		"map": {
			name:        "map",
			description: "TBA",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "TBA",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "TBA",
			callback:    commandExplore,
		},
	}
}

func commandExit(config commandConfig) ([]string, error) {
	_, err := fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil, err
}

func commandHelp(config commandConfig) ([]string, error) {
	usage := ""
	for _, command := range getRegistry() {
		usage += fmt.Sprintf("%s: %s\n", command.name, command.description)

	}
	_, err := fmt.Printf("Welcome to the Pokedex!\nUsage:\n\n%s", usage)
	return nil, err
}

func commandMap(config commandConfig) ([]string, error) {
	alt := ""
	if len(config.state) == 2 {
		alt = config.state[0]
	}
	return biMap(alt)
}
func commandMapB(config commandConfig) ([]string, error) {
	alt := ""
	if len(config.state) == 2 {
		alt = config.state[1]
	}
	return biMap(alt)
}

func commandExplore(config commandConfig) ([]string, error) {
	if len(config.args) == 0 {
		println("Command requires one argument")
		return nil, nil
	}
	areaTxt := config.args[0]
	area, err := pokeapi.GetArea("", areaTxt)
	if err != nil {
		return nil, err
	}
	for _, encounter := range area.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil, nil
}

func biMap(alt string) ([]string, error) {

	areas, err := pokeapi.GetAreas(alt)
	if err != nil {
		return nil, err
	}
	//fmt.Print("\n")
	for _, v := range areas.Results {

		fmt.Println(v.Name)
	}
	return []string{areas.Next, areas.Previous}, nil
}
