package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const promt = "pokedex"

func StartREPL(config *config) {
	StartMsg()
	input := bufio.NewScanner(os.Stdin)
	cli := GetCommands()

	for {
		fmt.Printf("%v > ", promt)
		input.Scan()

		words := cleanInput(input.Text())
		if len(words) == 0 {
			continue
		}

		command, ok := cli[words[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		err := command.Callback(config, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}

func StartMsg() {

	fmt.Println("Welcome to the world of Pokemons!")
	fmt.Println("Here you can: ")
	cli := GetCommands()

	for k, v := range cli {
		switch k {
		case "help":
			continue
		case "exit":
			continue
		case "mapb":
			continue
		case "map":
			continue
		default:
			fmt.Printf(" -%s\n", v.Description)

		}
	}
	fmt.Println()

}
