package main

import (
	"fmt"
	"os"

	"github.com/flying-house/pokedex/internal/pokeapi"
)

type command struct {
	name     string
	desc     string
	callback func(*config, []string) error
}

type config struct {
	nextURL   string
	prevURL   string
	apiClient *pokeapi.Client
}

var cfg config = config{
	nextURL:   "https://pokeapi.co/api/v2/location-area/",
	prevURL:   "",
	apiClient: pokeapi.NewClient(),
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
		"explore": {
			name:     "explore",
			desc:     "Explore a location area and see its Pokemon",
			callback: cmdExplore,
		},
	}
}

func cmdExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func cmdHelp(cfg *config, args []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.desc)
	}
	return nil
}

func cmdMap(cfg *config, args []string) error {
	if cfg.nextURL == "" {
		fmt.Println("No more locations to display")
		return nil
	}

	locations, err := cfg.apiClient.GetLocationAreas(cfg.nextURL)
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

func cmdMapBack(cfg *config, args []string) error {
	if cfg.prevURL == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := cfg.apiClient.GetLocationAreas(cfg.prevURL)
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

func cmdExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		fmt.Println("Please specify a location, e.g.:")
		fmt.Println("  explore pastoria-city-area")
		return nil
	}

	locationName := args[0]
	fmt.Printf("Exploring %s...\n", locationName)

	locationDetail, err := cfg.apiClient.GetLocationAreaDetail(locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationDetail.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}

	return nil
}
