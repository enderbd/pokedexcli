package main

import (
	"time"

	"github.com/enderbd/pokedexcli/internal/pokeapi"
)


func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config {
		pokeApiClient: pokeClient,
	}

	startRepl(cfg)
}

