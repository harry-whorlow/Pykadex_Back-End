package pokemonTypes

type PokemonBase struct {
	Attack       int `json:"attack"`
	Defense      int `json:"defense"`
	HealthPoints int `json:"healthPoints"`
	SpAttack     int `json:"spAttack"`
	SpDefense    int `json:"spDefense"`
	Speed        int `json:"speed"`
}

type PokemonName struct {
	Chinese  string `json:"chinese"`
	English  string `json:"english"`
	French   string `json:"french"`
	Japanese string `json:"japanese"`
}

type Pokemon struct {
	Base        PokemonBase `json:"base"`
	Description string      `json:"description"`
	Id          int         `json:"id"`
	Name        PokemonName `json:"name"`
	PokemonType string      `json:"pokemonType"`
	PokemonUrl  string      `json:"pokemonUrl"`
	Weakness    []string    `json:"weakness"`
}
