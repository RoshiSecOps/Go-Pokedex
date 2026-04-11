package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	formatText := strings.ToLower(text)
	endresult := strings.Fields(formatText)
	return endresult
}

func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return fmt.Errorf("")
}

func commandHelp() error {
	fmt.Print("Displays a help message\n")
	return fmt.Errorf("")
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Show help menu",
		callback:    commandHelp,
	},
}
