package pokeapi

type ResourceList struct {
	Name string `json:"name"`
	URL	 string `json:"url"`
}

type RespShallowLocations struct {
	Count 		int				`json:"count"`
	Next		*string			`json:"next"`
	Previous	*string			`json:"Previous"`
	Results		[]ResourceList	`json:"results"` 
}

type RespLocation struct {
	ID							int						`json:"id"`
	Name						string					`json:"name"`
	Game_index					int						`json:"game_index"`
	EncouterMethodRates			[]EncounterMethodRate	`json:"encounter_method_rates"`
	Location					Location				`json:"location"`
	Names						[]Names					`json:"names"`
	PokemonEncounters			[]PokemonEncounters		`json:"pokemon_encounters"`
}

type Location struct {
	Name string `json:"name"`
	URL	 string `json:"url"`
}

type Names struct {
	Language 	Language	`json:"language"`
	Name 		string		`json:"name"`
}

type Language struct {
	Name	string	`json:"name"`
	URL		string	`json:"url"`
}
