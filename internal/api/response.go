package api

type err string

func (e err) Error() string {
	return string(e)
}

const (
	ErrFetchResource = err("failed to fetch data")
	ErrDecode        = err("failed to decode json")
)

type LocationName struct {
	Name string `json:"name"`
}

type LocationsResponse struct {
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationName `json:"results"`
}

type PokemonName struct {
	Name string `json:"name"`
}

type ExploredPokemon struct {
	Pokemon PokemonName `json:"pokemon"`
}

type ExploreResponse struct {
	PokemonEncounters []ExploredPokemon `json:"pokemon_encounters"`
}
