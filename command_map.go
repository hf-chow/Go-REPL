package main

import (
	"errors"
	"fmt"
)

type Location struct {
	Count		int
	Next		string
	Previous	string
	Results		[]ResourceList
}

type ResourceList struct {
	Name		string
	URL			string
}

func commandMapf(cfg *config, param string) error {
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.prevLocationURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, param string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you are on the first page")
	}
	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationResp.Next
	cfg.prevLocationURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
