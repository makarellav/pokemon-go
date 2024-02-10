package repl

import (
	"bufio"
	"fmt"
	"github.com/makarellav/pokedex-repl/internal/api"
	"io"
	"log"
	"strings"
)

const prompt = "pokedex >> "

func Run(in io.Reader, out io.Writer, srv *api.PokemonService) {
	s := bufio.NewScanner(in)
	commands := getCommands()

	for {
		_, err := fmt.Fprint(out, prompt)

		if err != nil {
			log.Fatal(err)
		}

		scanned := s.Scan()

		if !scanned {
			return
		}

		args := cleanInput(s.Text())

		if len(args) == 0 {
			continue
		}

		command, ok := commands[args[0]]

		if !ok {
			fmt.Println("command not found")
			continue
		}

		if len(args) > 1 {
			args = args[1:]
		}

		err = command.callback(srv, args...)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}
