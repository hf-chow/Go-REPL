package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("please provide a pokemon name")
	}

	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...", name)

	exp := pokemon.BaseExperience
	success := catchSuccess(exp)
	if (success) {
		fmt.Printf("%s was caught!", name)
	} else {
		fmt.Printf("%s excaped!")
	}
	return nil
}

func catchSuccess(exp int) bool {
	success_prob := rand.Float32()
	exp_mod := float32(exp)/float32((len(strconv.Itoa(exp))+1))
	success_prob -= exp_mod
	if (success_prob > 0.5) {
		return true
	}
	return false
}
