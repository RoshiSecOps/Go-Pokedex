package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/RoshiSecOps/Go-Pokedex/internal/pokeapi"
)

type config struct {
	Next     *string
	Previous *string
}

func cleanInput(text string) []string {
	formatText := strings.ToLower(text)
	finalText := strings.Fields(formatText)
	return finalText
}

func commandExit(cfg *config) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	next, previous, err := pokeapi.GetLocations(url)
	if err != nil {
		return err
	}
	cfg.Next = &next
	cfg.Previous = &previous
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Displays a help message")
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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
