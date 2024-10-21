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
		fmt.Printf("-hp: %s", pokemon.Stats[0].Stat)
	} else {
		return errors.New("no such pokemon in the inventory")
	}
	return nil
} 
