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
		rawCommand := scanner.Text()
		commands := cleanInput(rawCommand)
		fmt.Printf("Your command was: %s\n", commands[0])
	}
}
