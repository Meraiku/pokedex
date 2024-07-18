package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/meraiku/pokedex/internal/world"
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

func StartMsg() *world.Player {
	data, err := os.ReadFile("./.save/player.txt")
	if err == nil {
		playerData := string(data)
		playerDataSlice := strings.Fields(playerData)
		name := playerDataSlice[1]
		age := playerDataSlice[3]
		ageNum, _ := strconv.Atoi(age)

		fmt.Printf("Welcome back %s!\n", name)
		return world.NewPlayer(name, ageNum)
	}

	fmt.Println("Welcome to the world of Pokemons!")
	fmt.Println("Here you can: ")
	cli := GetCommands()
	input := bufio.NewScanner(os.Stdin)

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
	fmt.Println("What's your name?")
	input.Scan()
	name := input.Text()

	fmt.Println("What's your age?")
	input.Scan()
	age := input.Text()
	ageNum, _ := strconv.Atoi(age)

	os.Mkdir("./.save", 0666)

	os.WriteFile("./.save/player.txt", []byte(fmt.Sprintf("name= %s\nage= %s", name, age)), 0666)
	return world.NewPlayer(name, ageNum)
}
