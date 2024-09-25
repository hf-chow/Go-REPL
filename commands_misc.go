package main

import (
	"fmt"
	"os"
)

func printPrompt() {
	fmt.Print("Pokedex > ")
}

func commandHelp(cfg *config, param string) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
    return nil
}

func commandExit(cfg *config, param string) error {
    os.Exit(0)
    return nil
}

func commandInvalid(cfg *config, param string) {
    fmt.Println("Invalid command")
}

