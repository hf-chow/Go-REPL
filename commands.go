package main

import (
	"fmt"
	"os"
)

func printPrompt() {
	fmt.Print(cliName, " > ")
}

func commandHelp(cfg *Config) error {
    fmt.Print("Welcome to the Pokedex!\n")
    fmt.Print("Usage: \n")
    fmt.Print()
    fmt.Print("help: Displays a help message\n")
    fmt.Print("exit: Exit the Pokedex\n")
    return nil
}

func commandExit(cfg *Config) error {
    os.Exit(0)
    return nil
}

func commandInvalid(cfg *Config) {
    fmt.Println("Invalid command")
}

