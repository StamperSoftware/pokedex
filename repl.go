package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func startRepl(config *Config) {

	userInput := []string{""}
	scan := bufio.NewScanner(os.Stdin)
	cmds := getCommands()	
	
	for userInput[0] != "exit" {
		fmt.Printf("Pokedex > ")
		scan.Scan()
		userInput = strings.Split(scan.Text(), " ")

		cmd, ok := cmds[userInput[0]]
		if ok {
			err := cmd.callback(config, userInput[1:])
			if err != nil {
				fmt.Printf("Error: %s\n",err)
			}
		} else {
			fmt.Println("Command not found")
		}	
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

