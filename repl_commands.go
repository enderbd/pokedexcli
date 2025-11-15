package main

import (
	"errors"
	"fmt"
	"os"

)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	helpDoc := `
Welcome to the Pokedex!
Usage:
 
`
	fmt.Print(helpDoc)
	for _, cmd := range commands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)

	}
	return nil
}

func commandMap(c *config) error {
	pokeReponse, err := c.pokeApiClient.GetLocations(c.nextLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = pokeReponse.Next
	c.prevLocationsURL = pokeReponse.Previous

	for _, loc := range pokeReponse.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapBack(c *config) error {
	if c.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	pokeReponse, err := c.pokeApiClient.GetLocations(c.prevLocationsURL)
	if err != nil {
		return err
	}

	c.nextLocationsURL = pokeReponse.Next
	c.prevLocationsURL = pokeReponse.Previous

	for _, loc := range pokeReponse.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandExplore(c *config) error {
	return nil
}

func commands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapBack,
		},
		"explore" : {
			name: "explore",
			description: "Explores the provided area",
			callback: commandExplore,
		},
	}
}
