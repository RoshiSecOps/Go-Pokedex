package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	fmt.Print("Welcome to the Pokedex!\n")
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input = cleanInput(scanner.Text())[0]
		value, ok := commands[input]
		if !ok {
			fmt.Print("Unknown command\n")
			continue
		}
		value.callback()
	}
}
