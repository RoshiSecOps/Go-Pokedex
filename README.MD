# Pokedex CLI

A command-line Pokedex built in Go. This tool allows users to explore the world of Pokemon by fetching real-time data from the PokéAPI.

## Features

- **Map Exploration**: Browse through different locations in the Pokemon world.
- **Pokemon Catching**: Attempt to catch Pokemon found in specific areas.
- **Pokedex Management**: View a list of all your caught Pokemon and their detailed stats.
- **Local Caching**: Optimized performance using a custom-built TTL (Time-To-Live) cache to reduce API calls.

## Installation

Ensure you have [Go](https://go.dev/dl/) installed on your machine.

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/pokedexcli.git
   cd pokedexcli

Build the application:
go build -o pokedex

2. Usage
Start the REPL by running the executable:

./pokedex

## Available Commands
- help: Displays a help message. - WIP
- exit: Exits the Pokedex.
- map: Displays the next 20 locations in the Pokemon world.
- mapb: Displays the previous 20 locations.
- explore <location_name>: Lists all Pokemon found in a specific area.
- catch <pokemon_name>: Attempts to catch a Pokemon and add it to your Pokedex.
- inspect <pokemon_name>: Shows stats, types, and info of a caught Pokemon.
- pokedex: Lists all the Pokemon you have caught.

## Technical Details
- Built using Go and the PokéAPI.
- Implements a REPL (Read-Eval-Print Loop) for user interaction.
- Uses a custom cache package to store API responses locally and improve speed.
