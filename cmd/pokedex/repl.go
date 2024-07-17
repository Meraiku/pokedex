package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const promt = "pokedex"

func StartREPL(config *config) {
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
		err := command.Callback(config)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
