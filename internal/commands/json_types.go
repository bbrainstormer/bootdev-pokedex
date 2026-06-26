package commands

// A list of types used to hold json data
// Many are incomplete; the extra fields are not used

type LocationArea struct {
	Id   int    `json:"id"`
	Name string `json:"string"`
	//GameIndex            int    `json:"game_index"`
	//EncounterMethodRates []EncounterMethodRate
	//Names []Name
	Encounters []Encounter `json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Height  int    `json:"height"`
	Weight  int    `json:"weight"`
	Species struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"species"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []PokemonType
}

type PokemonSpecies struct {
	CaptureRate int `json:"capture_rate"`
}

type PokemonType struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

type locationAreaOverview struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type locationResponse struct {
	Count    int                    `json:"count"`
	Next     string                 `json:"next"`
	Previous string                 `json:"previous"`
	Results  []locationAreaOverview `json:"results"`
}
