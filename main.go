package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	var answer string
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input = cleanInput(scanner.Text())[0]
		answer = "Your command was: " + input
		fmt.Println(answer)
	}
}
