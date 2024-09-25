package pokeapi

type RespEncounters struct {
	Count 		int		`json:"count"`
	Next		*string	`json:"next"`
	Previous	*string	`json:"Previous"`
	Results		[]struct {
		Name string `json:"name"`
		URL	 string `json:"url"`
	} `json: "results"`
}

type PokemonEncounters struct {
	Pokemon		Pokemon		`json:"pokemon"`
	-			string		`json:"version_details"`
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

