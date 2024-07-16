package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Location interface {
	NextMap() error
	PreviousMap() error
}

type PokeMap struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func NewPokeMap() *PokeMap {
	return &PokeMap{}
}

func (p *PokeMap) NextMap() error {

	err := getRequestToStruct(p, p.Next)
	if err != nil {
		return err
	}
	return nil
}

func (p *PokeMap) PreviousMap() error {
	if p.Previous == nil {
		fmt.Println("Cant")
		return fmt.Errorf("Can't go back in maps")
	}
	err := getRequestToStruct(p, *p.Previous)
	if err != nil {
		return err
	}
	return nil

}

func getRequestToStruct(p *PokeMap, direction string) error {
	if direction == "" {
		direction = "https://pokeapi.co/api/v2/location"
	}
	res, err := http.Get(direction)
	if err != nil {
		return errors.New("Cant acces map")
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return errors.New("Response failed with status code: %d and\nbody: %s\n")
	}
	if err != nil {
		return err
	}
	err = unmarhsal(body, p)
	if err != nil {
		return err
	}
	return nil
}

func unmarhsal(data []byte, p *PokeMap) error {
	err := json.Unmarshal(data, p)
	if err != nil {
		return err
	}
	return nil
}
