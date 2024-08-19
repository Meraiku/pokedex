package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/meraiku/pokedex/cmd/pokedex/structs"
	"github.com/meraiku/pokedex/internal/cache"
)

const startURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      cache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		cache: *cache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) LocationList(url *string) (*structs.PokeMap, error) {
	pokeMap := structs.PokeMap{}

	if url == nil {
		endpoint := "/location-area/"
		fullURL := startURL + endpoint
		url = &fullURL
	}

	cachedData, ok := c.cache.Get(*url)
	if ok {
		err := json.Unmarshal(cachedData, &pokeMap)
		if err != nil {
			return nil, err
		}
		return &pokeMap, nil
	}

	body, err := c.getRequest(*url)
	if err != nil {
		return nil, err
	}

	c.cache.Add(*url, body)

	err = json.Unmarshal(body, &pokeMap)
	if err != nil {
		return nil, err
	}
	return &pokeMap, nil
}

func (c *Client) PokemonList(areaName string) (*structs.LocationArea, error) {
	pokemon := structs.LocationArea{}

	endpoint := "/location-area/" + areaName
	url := startURL + endpoint

	cachedData, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return nil, err
		}
		return &pokemon, nil
	}

	body, err := c.getRequest(url)
	if err != nil {
		return nil, err
	}

	c.cache.Add(url, body)

	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return nil, err
	}
	return &pokemon, nil
}

func (c *Client) PokemonCatch(pokemonName string) (*structs.Pokemon, error) {
	pokemon := structs.Pokemon{}

	endpoint := "/pokemon/" + pokemonName
	url := startURL + endpoint

	cachedData, ok := c.cache.Get(url)
	if ok {
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return nil, err
		}
		return &pokemon, nil
	}

	body, err := c.getRequest(url)
	if err != nil {
		return nil, err
	}
	c.cache.Add(url, body)

	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return nil, err
	}
	return &pokemon, nil
}

func (c *Client) getRequest(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 399 {
		return nil, fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}
