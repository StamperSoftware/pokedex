package main

import (
	"strings"
	"bufio"
	"fmt"
	"os"
)

func startRepl(config *Config) {

	userInput := "" 
	scan := bufio.NewScanner(os.Stdin)
	cmds := getCommands()	
	
	for userInput != "exit" {
		fmt.Printf("Pokedex > ")
		scan.Scan()
		userInput = scan.Text()
		cmd, ok := cmds[userInput]
		if ok {
			err := cmd.callback(config)
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

