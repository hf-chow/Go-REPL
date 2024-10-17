package pokeapi

type Pokemon struct {
	ID						int					`json:"id"`
	Name					string				`json:"name"`
	BaseExperience			int					`json:"base_experience"`
	Height					int					`json:"height"`
	IsDefault				bool				`json:"is_default"`
	Order					int					`json:"order"`
	Weight					int					`json:"weight"`
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
