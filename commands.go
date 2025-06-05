package main

// excuse all the comments, my linter seems to believe they are required
// due to the capitalized names of the Command struct and I wanted to get
// rid of the warning flags

import (
	"fmt"
	"os"
)

// Command is... the commmand struct
type Command struct {
	name     string
	desc     string
	callback func() error
}

type config struct {
	nextURL string
	prevURL string
}

var cfg config = config{
	"ding",
	"dong",
}

// getCommands returns a map of commands
func getCommands() map[string]Command {
	return map[string]Command{
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

// cmdExit exits the Pokedex
func cmdExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// cmdHelp displays a help message
func cmdHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.desc)
	}
	return nil
}

func cmdMap() error {
	fmt.Println("Map not implemented yet")
	return nil
}

func cmdMapBack() error {
	fmt.Println("Back map not implemented yet")
	return nil
}
