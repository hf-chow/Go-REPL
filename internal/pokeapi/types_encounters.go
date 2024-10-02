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
	_			string		`json:"version_details"`
}

type EncounterMethodRate struct {
	EncounterMethod		EncounterMethod 	`json:"encounter_method"`
	_					string				`json:"version_details"`
						
}

type EncounterMethod struct {
	ID			int							`json:"id"`
	Name		string						`json:"name"`
	Order		int							`json:"order"`
	Names		[]Name						`json:"names"`
}

type EncounterConditionValue struct {
	ID 			int						`json:"id"`
	Name		string					`json:"name"`
	Condition	EncounterCondition		`json:"condition"`
	Names		[]Name					`json:"names"`
}

type EncounterCondition struct {
	ID			int							`json:"id"`
	Name		string						`json:"name"`
	Names		[]Name						`json:"names"`
	Values		[]EncounterConditionValue	`json:"values"`
}

type Name struct {
	Name	string	`json:"name"`
	Url		string	`json:"url"`
}

