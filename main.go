package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

			err := command.callback(&cfg)
			if err != nil {
				fmt.Println("command error: ", err)
			}
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(strings.TrimSpace(text)))
}
