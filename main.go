package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		rawInput := scanner.Text()
		commandNArgsTxt := cleanInput(rawInput)
		registry := getRegistry()
		command, exists := registry[commandNArgsTxt[0]]
		if exists {
			err := command.callback()
			if err != nil {
				println(err)
			}
			continue
		}
		println("Unknown command")
	}
}
