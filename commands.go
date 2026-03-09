package main

import (
	"fmt"
	"math/rand"
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
		"catch": {
			name:        "catch",
			description: "TBA",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "TBA",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "TBA",
			callback:    commandPokedex,
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

// Not the best placement, this file might get a bit messy.
var pokedex = make(map[string]pokeapi.Pokemon, 0)

func commandCatch(config commandConfig) ([]string, error) {
	if len(config.args) == 0 {
		println("Command requires one argument")
		return nil, nil
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", config.args[0])
	pokemon, err := pokeapi.GetPokemon("", config.args[0])
	if err != nil {
		return nil, err
	}
	baseXP := pokemon.BaseExperience
	chance := rand.Intn(baseXP/10 + 1)
	if chance != baseXP/10 {
		fmt.Printf("%s escaped!\n", config.args[0])
		return nil, nil
	}
	fmt.Printf("%s was caught!\n", config.args[0])
	pokedex[config.args[0]] = pokemon
	return nil, nil
}

func commandInspect(config commandConfig) ([]string, error) {
	if len(config.args) == 0 {
		println("Command requires one argument")
		return nil, nil
	}
	pokemon, exists := pokedex[config.args[0]]
	if !exists {
		fmt.Printf("%s was not in your pokedex", config.args[0])
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, v := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, v := range pokemon.Types {
		fmt.Printf(" - %s\n", v.Type.Name)
	}
	return nil, nil
}

func commandPokedex(config commandConfig) ([]string, error) {
	fmt.Println("Your Pokedex")
	for _, v := range pokedex {
		fmt.Printf(" - %s\n", v.Name)
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
