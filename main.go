package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var state []string = nil
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		rawInput := scanner.Text()
		commandNArgsTxt := cleanInput(rawInput)
		registry := getRegistry()
		command, exists := registry[commandNArgsTxt[0]]

		if exists {
			temp, err := command.callback(commandConfig{state: state, args: commandNArgsTxt[1:]})
			state = temp
			if err != nil {
				println(err)
			}
			continue
		}
		println("Unknown command")
	}
}
