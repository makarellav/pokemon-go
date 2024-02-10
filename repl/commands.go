package repl

import (
	"errors"
	"fmt"
	"github.com/makarellav/pokedex-repl/internal/api"
	"math/rand/v2"
	"os"
)

const minBaseExperience = 36

type cliCommand struct {
	name        string
	description string
	callback    func(service *api.PokemonService, params ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the Pokedex",
			callback:    commandExit,
		},
		"list": {
			name:        "list",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandList,
		},
		"list-back": {
			name:        "list-back",
			description: "Similar to 'list', however, it displays the previous 20 locations",
			callback:    commandListBack,
		},
		"explore": {
			name:        "explore <location-name>",
			description: "Displays the names of all Pokemons found in the specified location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon-name>",
			description: "Attempt to catch the specified Pokemon. The more experienced the Pokemon, the harder it is to catch!",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon-name>",
			description: "View details about a Pokemon in case you've already caught it",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View your Pokedex",
			callback:    commandPokedex,
		},
	}
}

func commandInspect(srv *api.PokemonService, params ...string) error {
	name := params[0]

	pokemon, ok := srv.PokemonCache.GetPokemon(name)

	if !ok {
		fmt.Println("You have not caught that pokemon")

		return nil
	}

	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("\t- %s\n", pokemonType.Type.Name)
	}

	return nil
}

func commandPokedex(srv *api.PokemonService, _ ...string) error {
	fmt.Println("Your Pokedex:")

	for k := range srv.PokemonCache.PokemonsData {
		pokemon, ok := srv.PokemonCache.GetPokemon(k)

		if !ok {
			fmt.Printf("failed to get %s\n", pokemon.Name)

			continue
		}

		fmt.Printf("\t- %s\n", pokemon.Name)
	}

	return nil
}

func commandCatch(srv *api.PokemonService, params ...string) error {
	name := params[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemonInfo, err := srv.CatchPokemon(name)

	if err != nil {
		return err
	}

	if rand.IntN(pokemonInfo.BaseExperience+1) <= minBaseExperience {
		fmt.Printf("%s was caught!\n", name)
		fmt.Println("You may now inspect it with the 'inspect' command")

		return nil
	}

	fmt.Printf("%s escaped!\n", name)

	return nil
}

func commandHelp(_ *api.PokemonService, _ ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commands := getCommands()

	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}

	fmt.Println()

	return nil
}

func commandExit(_ *api.PokemonService, _ ...string) error {
	os.Exit(0)

	return nil
}

func commandList(srv *api.PokemonService, _ ...string) error {
	locationsResp, err := srv.ListLocationAreas(srv.Config.Next)

	if err != nil {
		return err
	}

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	srv.Config.Next = locationsResp.Next
	srv.Config.Previous = locationsResp.Previous

	return nil
}

func commandListBack(srv *api.PokemonService, _ ...string) error {
	if srv.Config.Previous == "" {
		return errors.New("try to access invalid page")
	}

	locationsResp, err := srv.ListLocationAreas(srv.Config.Previous)

	if err != nil {
		return err
	}

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	srv.Config.Next = locationsResp.Next
	srv.Config.Previous = locationsResp.Previous

	return nil
}

func commandExplore(srv *api.PokemonService, params ...string) error {
	fmt.Printf("Exploring %s...\n", params[0])

	location, err := srv.ExploreLocation(params[0])

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, pokemonResp := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemonResp.Pokemon.Name)
	}

	return nil
}
