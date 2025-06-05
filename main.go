package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// is this where the URL goes? is this an "endpoint" or is there something else?
// are "startpoints" a thing? what does this URL do? what do I do with it?
const (
	apiURL = "https://pokeapi.co/api/v2/pokemon"
)

func main() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			input := scanner.Text()
			words := cleanInput(input)
			if len(words) == 0 {
				continue
			}

			cmd := words[0]
			command, exists := commands[cmd]

			if !exists {
				fmt.Println("Unknown command")
				continue
			}

			// why am I using a callback? why do some functions get called
			// directly and these do not? what the hell is even this
			err := command.callback()
			if err != nil {
				fmt.Println("command error: ", err)
			}
		}
	}
}

func getMapLocations() ([]string, error) {
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var data map[string]any

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return nil, nil //who fuckin knows how to return this thing, it's a map

}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}
