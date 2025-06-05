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

	return nil, nil

}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}
