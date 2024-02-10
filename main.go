package main

import (
	"github.com/makarellav/pokedex-repl/internal/api"
	"github.com/makarellav/pokedex-repl/internal/cache"
	"github.com/makarellav/pokedex-repl/repl"
	"os"
	"time"
)

func main() {
	pokemonService := api.NewService(api.NewConfig(), cache.NewRequestCache(time.Minute), cache.NewPokemonCache())

	repl.Run(os.Stdin, os.Stdout, pokemonService)
}
