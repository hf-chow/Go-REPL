package main 

import (
)

var cliName string = "Pokedex"

type Config struct {
	Next		string
	Previous	string
}

func main(){
	cfg := &Config {
		Next: 		"",
		Previous: 	"",
	}

	repl(cfg)
}

