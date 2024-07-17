package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/meraiku/pokedex/internal/pokeapi"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *config) error
}

func GetCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Dispays list of 20 maps",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays list of 20 previous maps",
			Callback:    commandMapb,
		},
	}
}

func commandHelp(c *config) error {

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

func commandExit(c *config) error {
	os.Exit(0)
	return nil
}

func commandMap(c *config) error {
	locations, err := c.pokeAPIClient.LocationList(c.nextLocationAreaURL)
	if err != nil {
		return err
	}
	c.nextLocationAreaURL = locations.Next
	c.previousLocationAreaURL = locations.Previous

	printLocations(locations)
	return nil
}

func commandMapb(c *config) error {
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
	fmt.Println("Locatinos: ")
	for _, v := range locations.Results {
		fmt.Printf("  - %s\n", v.Name)
	}
}
