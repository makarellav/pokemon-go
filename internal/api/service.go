package api

import (
	"encoding/json"
	"github.com/makarellav/pokedex-repl/internal/cache"
	"github.com/makarellav/pokedex-repl/internal/model"
	"io"
	"net/http"
)

type PokemonService struct {
	Config       *Config
	RequestCache *cache.RequestCache
	PokemonCache *cache.PokemonCache
}

func NewService(config *Config, cache *cache.RequestCache, pokemonCache *cache.PokemonCache) *PokemonService {
	return &PokemonService{
		Config:       config,
		RequestCache: cache,
		PokemonCache: pokemonCache,
	}
}

func (ps *PokemonService) ListLocationAreas(path string) (*LocationsResponse, error) {
	url := baseURL + "/location-area"

	if path != "" {
		url = path
	}

	return getResource[LocationsResponse](url, ps.RequestCache)
}

func (ps *PokemonService) ExploreLocation(locationName string) (*ExploreResponse, error) {
	return getResource[ExploreResponse](baseURL+"/location-area/"+locationName, ps.RequestCache)
}

func (ps *PokemonService) CatchPokemon(pokemonName string) (*model.Pokemon, error) {
	return getResource[model.Pokemon](baseURL+"/pokemon/"+pokemonName, ps.PokemonCache)
}

func getResource[T any](url string, cache cache.Cache) (*T, error) {
	var response T

	if data, ok := cache.Get(url); ok {
		if err := json.Unmarshal(data, &response); err != nil {
			return nil, ErrDecode
		}

		return &response, nil
	}

	resp, err := http.Get(url)

	defer resp.Body.Close()

	if err != nil {
		return nil, ErrFetchResource
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &response)

	if err != nil {
		return nil, ErrDecode
	}

	cache.Add(url, data)

	return &response, nil
}
