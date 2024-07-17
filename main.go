package main 

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

var cliName string = "Pokedex"

type cliCommand struct {
    name        string
    description string
    callback    func() error
}

func commandHelp() error {
    print("Welcome to the Pokedex!\n")
    print("Usage: \n")
    print()
    print("help: Displays a help message\n")
    print("exit: Exit the Pokedex\n")
    return nil
}

func commandExit() error {
    os.Exit(0)
    return nil
}

func printPrompt() {
    fmt.Print(cliName, " > ")
}

func commandInvalid() {
    fmt.Println("Invalid command")
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
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


func main(){

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

