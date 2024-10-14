package pokeapi

type RespPokemon struct {
	ID						int					`json:"id"`
	Name					string				`json:"name"`
	BaseExperience			int					`json:"base_experience"`
	Height					int					`json:"height"`
	IsDefault				bool				`json:"is_default"`
	Order					int					`json:"order"`
	Weight					int					`json:"weight"`
	Abilities				[]string			`json:"abilities"`
	Forms					[]string			`json:"forms"`
	GameIndices				[]string			`json:game_indices"`
	HeldItems				[]string			`json:held_items"`
	LocationAreaEncounters	string				`json:"location_area_encounters"`
	Moves					[]string			`json:"moves"`
	_						[]string			`json:"past_types"`
	_						string				`json:"sprites"`
	_						string				`json:"cries"`
	_						string				`json:"species"`
	stats					[]PokemonStat		`json:"stats"`
	types					[]PokemonType		`json:"types"`
}

type PokemonStat struct {
	Name		string	`json:"stat"`
	Effort		int		`json:"effort"`
	BaseStat	int		`json:"base_stat"`
}

type PokemonType struct {
	Slot		int		`json:"slot"`
	Type		string	`json:"type"`
}
