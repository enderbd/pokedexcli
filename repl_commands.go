package main

import (
	"errors"
	"fmt"
	"os"
	"math/rand"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func commandExit(c *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, args []string) error {
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

func commandMap(c *config, args []string) error {
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

func commandMapBack(c *config, args []string) error {
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

func commandExplore(c *config, args []string) error {
	if len(args) < 1  {
		return  errors.New("Area to explore not provided!")
	}

	area := args[0]
	areaInfo, err := c.pokeApiClient.GetPokemonEncounters(area)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", areaInfo.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range areaInfo.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}

func commandCatch(c* config, args []string) error {
	if len(args) < 1 {
		return errors.New("Need a pokemong name to catch!")
	}
	pokemonName := args[0]

	pokemonInfo, err := c.pokeApiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}
	pokemoneBaseExp := pokemonInfo.BaseExperience
	rngCatch := rand.Intn(pokemoneBaseExp)


	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonInfo.Name)
	if rngCatch > 40 {
		fmt.Printf("%s escaped!\n", pokemonInfo.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonInfo.Name)
	c.caughtPokemon[pokemonInfo.Name] = pokemonInfo

	return nil
}

func commandInspect(c* config, args []string) error {
	if len(args) < 1 {
		return errors.New("Need a pokemon name!")
	}
	pokemonName := args[0]

	pokemonInfo, ok := c.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("Need to catch it first before we can see the stats")
	}

	fmt.Printf("Name: %s\n", pokemonInfo.Name)
	fmt.Printf("Height: %v\n", pokemonInfo.Height)
	fmt.Printf("Weight: %v\n", pokemonInfo.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemonInfo.Stats {
		fmt.Printf("  -%s:%v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemonInfo.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

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
			name: "explore <area_name>",
			description: "Explores the provided area",
			callback: commandExplore,
		},
		"catch" : {
			name: "catch <pokemon_name>",
			description: "Attempt to catches the pokemon !",
			callback: commandCatch,
		},
		"inspect" : {
			name: "inspect <pokemon_name>",
			description: "If the <pokemon-name> was caught, display it's stats !",
			callback: commandInspect,
		},

}
}
