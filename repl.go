package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/RoshiSecOps/Go-Pokedex/internal/pokeapi"
	"github.com/RoshiSecOps/Go-Pokedex/internal/pokecache"
)

type config struct {
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
	Pokedex  map[string]pokeapi.Pokemon
}

func cleanInput(text string) []string {
	formatText := strings.ToLower(text)
	finalText := strings.Fields(formatText)
	return finalText
}

func commandExit(cfg *config, args ...string) error {
	fmt.Print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}
func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a location name")
	}
	pokeapi.GetPokemons(args[0], cfg.Cache)
	return nil
}

func commandMapBack(cfg *config, args ...string) error {
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
func commandCatch(cfg *config, args ...string) error {
	var monsterTier int
	targetName := strings.ToLower(args[0])
	fmt.Printf("Throwing a Pokeball at %v...\n", targetName)
	diceRoll := rand.Intn(10)
	targetPokemon, err := pokeapi.GetPokemonStats(targetName, cfg.Cache)
	if err != nil {
		return err
	}
	if targetPokemon.BaseExperience < 300 {
		monsterTier = 3
	} else if targetPokemon.BaseExperience < 800 {
		monsterTier = 5
	} else {
		monsterTier = 8
	}
	if diceRoll > monsterTier {
		fmt.Printf("%s was caught!\n", targetName)
		cfg.Pokedex[targetName] = targetPokemon
		return nil
	} else {
		fmt.Printf("%s escaped!\n", targetName)
		return nil
	}
}

func commandMap(cfg *config, args ...string) error {
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

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("Displays a help message")
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"catch": {
		name:        "catch pokemon",
		description: "Attempt to catch a pokemon",
		callback:    commandCatch,
	},
	"explore": {
		name:        "explore",
		description: "Display Pokemons from a give location",
		callback:    commandExplore,
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
