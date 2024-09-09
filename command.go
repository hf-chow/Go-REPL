package main

import (
	"fmt"
	"os"
)

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
