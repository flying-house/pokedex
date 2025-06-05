package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type command struct {
	name     string
	desc     string
	callback func(*config) error
}

type config struct {
	nextURL string
	prevURL string
}

// LocationArea comment for linter
type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LocationResponse comment for linter
type LocationResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

var cfg config = config{
	nextURL: "https://pokeapi.co/api/v2/location-area/",
	prevURL: "",
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			name:     "help",
			desc:     "Displays a help message",
			callback: cmdHelp,
		},
		"exit": {
			name:     "exit",
			desc:     "Exit the Pokedex",
			callback: cmdExit,
		},
		"map": {
			name:     "map",
			desc:     "Displays twenty locations at a time, advancing with each call",
			callback: cmdMap,
		},
		"mapb": {
			name:     "mapb",
			desc:     "Displays the previous twenty locations",
			callback: cmdMapBack,
		},
	}
}

func getLocations(url string) (*LocationResponse, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var locationResponse LocationResponse
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return nil, err
	}

	return &locationResponse, nil
}

func cmdExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cmdHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.desc)
	}
	return nil
}

func cmdMap(cfg *config) error {
	if cfg.nextURL == "" {
		fmt.Println("No more locations to display")
		return nil
	}

	locations, err := getLocations(cfg.nextURL)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	if locations.Next != nil {
		cfg.nextURL = *locations.Next
	} else {
		cfg.nextURL = ""
	}

	if locations.Previous != nil {
		cfg.prevURL = *locations.Previous
	} else {
		cfg.prevURL = ""
	}

	return nil
}

func cmdMapBack(cfg *config) error {
	if cfg.prevURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := getLocations(cfg.prevURL)
	if err != nil {
		return err
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	if locations.Next != nil {
		cfg.nextURL = *locations.Next
	} else {
		cfg.nextURL = ""
	}

	if locations.Previous != nil {
		cfg.prevURL = *locations.Previous
	} else {
		cfg.prevURL = ""
	}

	return nil
}
