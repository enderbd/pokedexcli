package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
	"github.com/enderbd/pokedexcli/internal/pokecache"
)

// Base URL for the Poke API
const baseURL =  "https://pokeapi.co/api/v2"

type PokeResponse struct {
	Count    int     `json:"coun"`
	Next     *string `json:"next"`
	Previous *string `json:"previus"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokeAreaInfo struct {
	Name  string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// Client struct
type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

// NewClient 
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval)	,
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}


// get the locations form the provided URL 
func (c* Client) GetLocations(fromURL *string) (PokeResponse, error) {
	// first call will be for the start point,
	url := baseURL + "/location-area"
	if fromURL != nil {
		url = *fromURL
	}
	
	if val, ok := c.cache.Get(url); ok {
		pokeResonse := PokeResponse{}
		err := json.Unmarshal(val, &pokeResonse)
		if err != nil {
			return PokeResponse{}, err
		}
		return pokeResonse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeResponse{}, err
	}

	resp, err := c.httpClient.Do(req)		
	if err != nil {

		return PokeResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeResponse{}, err
	}

	pokeResponse := PokeResponse{}
	err = json.Unmarshal(data, &pokeResponse)
	if err != nil {
		return PokeResponse{}, nil
	}
	
	c.cache.Add(url, data)
	return pokeResponse, nil
}

func (c* Client) GetPokemonEncounters(area string) (PokeAreaInfo, error) {
	url := baseURL + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		pokeAreaInfo := PokeAreaInfo{}
		err := json.Unmarshal(val, &pokeAreaInfo)
		if err != nil {
			return PokeAreaInfo{}, err
		}
		return pokeAreaInfo, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeAreaInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeAreaInfo{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeAreaInfo{}, err
	}

	pokeAreaInfo := PokeAreaInfo{}
	err = json.Unmarshal(dat, &pokeAreaInfo)
	if err != nil {
		return PokeAreaInfo{}, err
	}

	c.cache.Add(url, dat)

	
	return pokeAreaInfo, nil
}
