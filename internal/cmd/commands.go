package cmd

import (
	"fmt"
	"os"

	"github.com/meraiku/pokedex/internal/pokeapi"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(c *pokeapi.PokeMap) error
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

func commandHelp(c *pokeapi.PokeMap) error {

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

func commandExit(c *pokeapi.PokeMap) error {
	os.Exit(0)
	return nil
}

func commandMap(c *pokeapi.PokeMap) error {
	c.NextMap()

	for _, v := range c.Results {
		fmt.Println(v.Name)
	}
	return nil
}

func commandMapb(c *pokeapi.PokeMap) error {
	if err := c.PreviousMap(); err != nil {
		return err
	}
	for _, v := range c.Results {
		fmt.Println(v.Name)
	}
	return nil
}
