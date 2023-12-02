package main

import (
	"fmt"
	"os"
)

func helpCommand() error{
	commands := getCommands()

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func exitCommand() error {
	os.Exit(0)
	return nil
}

func pokemonCommand() error {
	
}