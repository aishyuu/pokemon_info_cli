package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type command struct {
	name string
	description string
	callback func() error
}

func cliStart() {
	reader := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("> ")
		reader.Scan()

		cleanedInput := cleanInput(reader.Text())

		command, ok := commands[cleanedInput]
		if ok {
			command.callback()
			continue
		}
		fmt.Println("Command does not exist, try again.")
	}
}

func cleanInput(input string) string {
	return strings.ToLower(input)
}

func getCommands() map[string]command {
	return map[string]command {
		"help" : {
			name: "help",
			description: "Displays a help message.",
			callback: helpCommand,
		},
		"exit" : {
			name: "exit",
			description: "Exit the application.",
			callback: exitCommand,
		},
	}
}