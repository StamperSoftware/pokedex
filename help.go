package main

import (
	"fmt"
)

func commandHelp(c *Config, params []string) error {
	fmt.Println("Commands:")
	for _, cmd := range getCommands() {
		fmt.Printf("Command %s:\n--------------\n%s\n\n", cmd.name, cmd.description)
	}
	return nil
}


