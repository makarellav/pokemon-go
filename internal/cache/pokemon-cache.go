package cache

import (
	"encoding/json"
	"github.com/makarellav/pokedex-repl/internal/model"
	"path"
)

type PokemonCache struct {
	PokemonsData map[string][]byte
}

func NewPokemonCache() *PokemonCache {
	return &PokemonCache{PokemonsData: make(map[string][]byte)}
}

func (pc *PokemonCache) Add(key string, data []byte) {
	pc.PokemonsData[path.Base(key)] = data
}

func (pc *PokemonCache) Get(key string) ([]byte, bool) {
	pokemon, ok := pc.PokemonsData[key]

	return pokemon, ok
}

func (pc *PokemonCache) GetPokemon(key string) (*model.Pokemon, bool) {
	data, ok := pc.Get(key)

	if !ok {
		return nil, ok
	}

	var pokemon model.Pokemon

	err := json.Unmarshal(data, &pokemon)

	if err != nil {
		return nil, false
	}

	return &pokemon, true
}
