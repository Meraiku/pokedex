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
type LocationArea struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height    int `json:"height"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name          string `json:"name"`
	Order         int    `json:"order"`
	PastAbilities []any  `json:"past_abilities"`
	PastTypes     []any  `json:"past_types"`
	Species       struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
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

func (c *Client) PokemonList(areaName string) (*LocationArea, error) {
	pokemon := LocationArea{}

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

func (c *Client) PokemonCatch(pokemonName string) (*Pokemon, error) {
	pokemon := Pokemon{}

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
