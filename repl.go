package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/RoshiSecOps/Go-Pokedex/internal/pokeapi"
	"github.com/RoshiSecOps/Go-Pokedex/internal/pokecache"
)

type config struct {
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
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

func commandMapBack(cfg *config) error {
	var url string
	if cfg.Previous == nil || *cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	url = *cfg.Previous
	next, previous, err := pokeapi.GetLocations(url, cfg.Cache)
	if err != nil {
		return err
	}
	cfg.Next = &next
	cfg.Previous = &previous
	return nil
}

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil {
		url = *cfg.Next
	}
	next, previous, err := pokeapi.GetLocations(url, cfg.Cache)
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
	"mapb": {
		name: "map",
		description: `Display previous 20 location areas in the pokemon world.
		Subsequent calls display even older ones, if not on first page.`,
		callback: commandMapBack,
	},
	"help": {
		name:        "help",
		description: "Show help menu",
		callback:    commandHelp,
	},
}
