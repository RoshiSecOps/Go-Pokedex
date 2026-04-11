package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	formatText := strings.ToLower(text)
	finalText := strings.Fields(formatText)
	return finalText
}

func commandMap() error {
	return nil
}
func commandExit() error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Displays a help message")
	return nil
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
	"map": {
		name: "map",
		description: `Display 20 location areas in the pokemon world.
		Subsequent calls display the next 20.`,
		callback: commandMap,
	},
	"help": {
		name:        "help",
		description: "Show help menu",
		callback:    commandHelp,
	},
}
