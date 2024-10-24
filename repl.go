package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hf-chow/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient	pokeapi.Client
	nextLocationURL	*string
	prevLocationURL	*string
	pokemonCaught	map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name 		string
	description	string
	callback	func(c *config, args ...string) error
}

func startRepl(cfg *config) {
    reader := bufio.NewScanner(os.Stdin)
	cfg.pokemonCaught = make(map[string]pokeapi.Pokemon)

    for {
        printPrompt()
		reader.Scan()

        words := cleanInput(reader.Text())
        if len(words) == 0 {
            continue
        }

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists{
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
			} else {
				fmt.Println("Unknown command")
				continue
			}
	}
}


func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
		"catch": {
			name: 			"catch",
			description: 	"Attempt to catch a pokemon",
			callback:		commandCatch,
		},
		"inspect": {
			name: 			"inspect",
			description: 	"Inpect a caught pokemon",
			callback:		commandInspect,
		},
        "explore": {
            name:           "explore",
            description:    "Displays the pokemons in this location",
            callback:       commandExplore,
        },
        "map": {
            name:           "map",
            description:    "Displays the first 20 or next 20 locations",
            callback:       commandMapf,
        },
        "mapb": {
            name:           "map back",
            description:    "Displays the previous 20 locations",
            callback:       commandMapb,
        },
        "pokedex": {
            name:           "pokedex",
            description:    "Displays all the Pokemon caught",
            callback:       commandPokedex,
        },
        "help": {
            name:           "help",
            description:    "Displays a help message",
            callback:       commandHelp,
        },
        "exit": {
            name:           "exit",
            description:    "Exit the Pokedex",
            callback:       commandExit,
        },
    }
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}
