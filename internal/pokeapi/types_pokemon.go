package pokeapi

type Pokemon struct {
	ID						int					`json:"id"`
	Name					string				`json:"name"`
	BaseExperience			int					`json:"base_experience"`
	Height					int					`json:"height"`
	IsDefault				bool				`json:"is_default"`
	Order					int					`json:"order"`
	Weight					int					`json:"weight"`
	Stats					[]PokemonStats		`json:"stats"`
	Types					[]PokemonType		`json:"types"`
}

type PokemonStats struct {
	Stat		Stat		`json:"stat"`
	Effort		int			`json:"effort"`
	BaseStat	int			`json:"base_stat"`
}

type Stat struct {
	ID 				int			`json:"id"`
	Name			string		`json:"name"`
	GameIndex		int			`json:"game_index"`
	IsBattleOnly	bool		`json:"is_battle_only"`
}

type PokemonType struct {
	Slot		int		`json:"slot"`
	Type		Type	`json:"type"`
}

type Type struct {
	Name		string  `json:"name"`
}
