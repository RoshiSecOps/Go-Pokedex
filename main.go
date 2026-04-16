package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/RoshiSecOps/Go-Pokedex/internal/pokeapi"
	"github.com/RoshiSecOps/Go-Pokedex/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Welcome to the Pokedex!\n")
	cache := pokecache.NewCache(10 * time.Second)
	pokedex := make(map[string]pokeapi.Pokemon)
	cfg := config{}
	cfg.Cache = cache
	cfg.Pokedex = pokedex
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		input := words[0]
		command, ok := commands[input]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		if len(words) > 1 {
			if err := command.callback(&cfg, words[1]); err != nil {
				fmt.Println(err)
			}
		} else {
			if err := command.callback(&cfg); err != nil {
				fmt.Println(err)
			}
		}
	}
}
