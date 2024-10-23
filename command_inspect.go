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
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")
		fmt.Printf("  -hp: %d\n", pokemon.Stats[0].BaseStat)
		fmt.Printf("  -attack: %d\n", pokemon.Stats[1].BaseStat)
		fmt.Printf("  -defense: %d\n", pokemon.Stats[2].BaseStat)
		fmt.Printf("  -special-attack: %d\n", pokemon.Stats[3].BaseStat)
		fmt.Printf("  -special-defense: %d\n", pokemon.Stats[4].BaseStat)
		fmt.Printf("  -speed: %d\n", pokemon.Stats[5].BaseStat)
		fmt.Println("Types:")
		for i:=0; i<len(pokemon.Types); i++ {
			fmt.Printf("  - %s\n", pokemon.Types[i].Type.Name)
		}
	} else {
		return errors.New("no such pokemon in the inventory")
	}
	return nil
} 
