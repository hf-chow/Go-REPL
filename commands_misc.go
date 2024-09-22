package main

import (
	"fmt"
	"os"
)

func printPrompt() {
	fmt.Println(cliName, " > ")
}

func commandHelp(cfg *Config) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:")
    fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.cmd.description)
	}
	fmt.Println()
    return nil
}

func commandExit(cfg *Config) error {
    os.Exit(0)
    return nil
}

func commandInvalid(cfg *Config) {
    fmt.Println("Invalid command")
}

