package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const startURL = "https://pokeapi.co/api/v2"

type Client struct {
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

func NewClient() *Client {
	return &Client{
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

	err = json.Unmarshal(body, &pokeMap)
	if err != nil {
		return nil, err
	}
	return &pokeMap, nil
}
