package main

import (
	"fmt"
	"os"
)

// Command is a basic struct for definitions
type Command struct {
	name     string
	desc     string
	callback func() error
}

// getCommands returns a map of commands with their name, description, and callback function
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
