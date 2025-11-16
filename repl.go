package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/enderbd/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeApiClient pokeapi.Client
	prevLocationsURL *string
	nextLocationsURL *string
	caughtPokemon map[string]pokeapi.PokemonInfo
}


func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex > "
	var userInput []string
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}


		userInput = cleanInput(scanner.Text())

		if len(userInput) == 0 {
			continue
		}
		commandName := userInput[0]
		var args []string
		if len(userInput) > 1 {
			args = userInput[1:]
		}
		
		command, exists := commands()[commandName]
		if exists {
			err := command.callback(cfg, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	var result []string
	result = strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return result
}
