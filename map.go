package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func commandMap(cfg *Config) error {

	var location Location

	resp, err := http.Get(cfg.Next)
	if err != nil {
		return fmt.Errorf("error when getting locations: %v", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return fmt.Errorf("error when decoding: %v", err)
	}

	for _, loc := range location.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	cfg.Next = location.Next
	cfg.Previous = location.Previous

	return nil
}

func commandMapBack(cfg *Config) error {

	var location Location

	if cfg.Previous == "" {
		fmt.Printf("You are in the first page of the map!")
		return nil
	}

	resp, err := http.Get(cfg.Previous)
	if err != nil {
		return fmt.Errorf("error when getting locations: %v", err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&location)
	if err != nil {
		return fmt.Errorf("error when decoding: %v", err)
	}

	for _, loc := range location.Results {
		fmt.Printf("%s\n", loc.Name)
	}

	cfg.Next = location.Next
	cfg.Previous = location.Previous

	return nil
}
