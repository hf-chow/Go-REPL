package main

import (
	"fmt"
	"errors"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a pokemon name")
	}
	name := args[0]
	pokemon, ok := cfg.pokemonCaught[name]
	if ok {
		fmt.Printf("-hp: %d\n", pokemon.Stats[0].BaseStat)
	} else {
		return errors.New("no such pokemon in the inventory")
	}
	return nil
} 
