package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func repl(cfg *Config) {
    scanner := bufio.NewScanner(os.Stdin)
    printPrompt()

    for scanner.Scan() {

        printPrompt()

        words := cleanInput(scanner.Text())

        if len (words) == 0 {
            continue
        }

        commandName := words[0]

        command, exists := getCommands()[commandName]
        if exists {
            err := command.callback(cfg)
            if err != nil {
                fmt.Println(err)
            }
            continue
        } else {
            commandInvalid(cfg)
			printPrompt()
            continue
        }
    }
}

type cliCommand struct {
    name        string
    description string
    callback    func(*Config) error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "map": {
            name:           "map",
            description:    "Displays the first 20 or next 20 locations",
            callback:       commandMap,
        },
        "mapb": {
            name:           "map back",
            description:    "Displays the previous 20 locations",
            callback:       commandMapBack,
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
