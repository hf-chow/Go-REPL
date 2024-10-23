package main 

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.pokemonCaught) == 0 {
		return errors.New("no pokemon in your inventory")
	} else {
		fmt.Println("Your Pokedex:")
		for k, _ := range cfg.pokemonCaught {
			fmt.Printf(" - %s\n", k)
		}
	}
	return nil
}
