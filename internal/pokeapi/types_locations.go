package pokeapi

type Location struct {
	Count		int
	Next		string
	Previous	string
	Results		[]ResourceList
}

type ResourceList struct {
	Name string `json:"name"`
	URL	 string `json:"url"`
}

type RespShallowLocations struct {
	Count 		int				`json:"count"`
	Next		*string			`json:"next"`
	Previous	*string			`json:"Previous"`
	Results		ResourceList	`json:"results"` 
}

type RespLocationArea struct {
	ID							int						`json:"id"`
	Name						string					`json:"name"`
	Game_index					int						`json:"game_index"`
	EncouterMethodRates			[]EncounterMethodRate	`json:"encounter_method_rates"`
	Location					Location				`json:"location"`
	Names						[]string				`json:"names"`
	PokemonEncounters			[]PokemonEncounters		`json:"pokemon_encounters"`
}

