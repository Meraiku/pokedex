package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/meraiku/pokedex/internal/cache"
)

const startURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      cache.Cache
	httpClient http.Client
}

type PokeMap struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func NewClient(cacheInterval time.Duration) *Client {
	return &Client{
		cache: *cache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) LocationList(url *string) (*PokeMap, error) {
	pokeMap := PokeMap{}

	if url == nil {
		endpoint := "/location/"
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

	request, err := http.NewRequest("GET", *url, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 399 {
		return nil, fmt.Errorf("reqponse failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
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
