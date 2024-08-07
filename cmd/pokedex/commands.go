package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/meraiku/pokedex/internal/pokeapi"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*config, ...string) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "List of next 20 location-areas",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "List of previous 20 location-areas",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore {area name}",
			Description: "Explore location-area to find pokemons",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch {pokemon name}",
			Description: "Throw pokeball to catch pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect {pokemon name}",
			Description: "Inspect pokemon from your pokedex",
			Callback:    commandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Open your Pokedex",
			Callback:    commandPokedex,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
	}
}

func commandHelp(c *config, args ...string) error {

	fmt.Println("")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println("")

	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println("")

	return nil
}

func commandExit(c *config, args ...string) error {
	os.Exit(0)
	return nil
}

func commandMap(c *config, args ...string) error {

	locations, err := c.pokeAPIClient.LocationList(c.nextLocationAreaURL)
	if err != nil {
		return err
	}

	c.nextLocationAreaURL = locations.Next
	c.previousLocationAreaURL = locations.Previous

	printLocations(locations)
	return nil

}

func commandMapb(c *config, args ...string) error {
	if c.previousLocationAreaURL == nil {
		return errors.New("can't go back in maps")
	}

	locations, err := c.pokeAPIClient.LocationList(c.previousLocationAreaURL)
	if err != nil {
		return err
	}

	c.nextLocationAreaURL = locations.Next
	c.previousLocationAreaURL = locations.Previous

	printLocations(locations)
	return nil
}

func printLocations(locations *pokeapi.PokeMap) {
	fmt.Println("Locatinos areas: ")
	for _, v := range locations.Results {
		fmt.Printf("  - %s\n", v.Name)
	}
}

func commandExplore(c *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	areaName := args[0]

	pokemons, err := c.pokeAPIClient.PokemonList(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", areaName)
	time.Sleep(time.Second)
	fmt.Println("Pokemons found: ")
	for _, v := range pokemons.PokemonEncounters {
		fmt.Printf("  - %s\n", v.Pokemon.Name)
		time.Sleep(time.Millisecond * 50)
	}

	return nil

}

func commandCatch(c *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemons, err := c.pokeAPIClient.PokemonCatch(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	time.Sleep(time.Second * 2)
	chacnceToCatch := rand.Intn(pokemons.BaseExperience)
	if chacnceToCatch > pokemons.BaseExperience/2 {

		fmt.Printf("%s was caught!\n", pokemonName)
		c.player.Pokedex.Pokedex[pokemonName] = *pokemons
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemonName)

	return nil

}

func commandInspect(c *config, args ...string) error {

	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonInfo, ok := c.player.Pokedex.Pokedex[args[0]]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\nHeight: %v\nWeight: %v\nStats: \n", pokemonInfo.Name, pokemonInfo.Height, pokemonInfo.Weight)

	for _, v := range pokemonInfo.Stats {
		fmt.Printf("  -%s: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")

	for _, v := range pokemonInfo.Types {
		fmt.Printf("  - %s\n", v.Type.Name)
	}

	return nil
}

func commandPokedex(c *config, args ...string) error {
	fmt.Println("Your Pokedex: ")

	for k := range c.player.Pokedex.Pokedex {
		fmt.Printf("  - %s\n", k)
	}
	return nil
}
