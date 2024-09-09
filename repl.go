package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repl() {
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
            err := command.callback()
            if err != nil {
                fmt.Println(err)
            }
            continue
        } else {
            commandInvalid()
            continue
        }
    }
}

type cliCommand struct {
    name        string
    description string
    callback    func() error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
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
