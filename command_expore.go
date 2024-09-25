package main

import (
	"errors"
)

type LocationArea struct {
	ID							int
	Name						string
	Game_index					int
	EncouterMethodRates		[]EncounterMethodRate
	Location					Location
	Names						[]string
	PokemonEncounters			[]PokemonEncounters
}

type EncounterMethodRate struct {
	EncounterMethod		EncounterMethod
	_					
}

type EncounterVersionDetail struct {
	Rate			int
	Version			EncounterVersionDetails
}


type EncounterMethod struct {
	ID			int
	Name		string
	Names		[]string
	Values		[]EncounterConditionValue
}

type EncounterConditionValue struct {
	ID 			int
	Name		string
	Condition	EncounterCondition
	Names		[]string
}

type EncounterCondition struct {
	ID			int
	Name		string
	Names		[]string
	Values		[]EncounterConditionValue
}




func commandExplore(cfg *config) error {
}
