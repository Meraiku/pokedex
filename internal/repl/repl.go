package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/meraiku/pokedex/internal/cmd"
	"github.com/meraiku/pokedex/internal/pokeapi"
)

const promt = "pokedex"

func StartREPL() {
	input := bufio.NewScanner(os.Stdin)
	cli := cmd.GetCommands()
	config := pokeapi.NewPokeMap()

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
