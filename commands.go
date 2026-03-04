package main

type cliCommand struct {
	name string
	description string
	callback func(c *Config, params []string) error
}

func getCommands() map[string]cliCommand{
	
	return map[string]cliCommand{
		"exit": {
			name:"exit",
			description:"closes program with exit code 0",
			callback:commandExit,
		},
		"help": {
			name:"help",
			description:"Prints help menu",
			callback:commandHelp,
		},
		"map": {
			name:"map",
			description:"Returns next list of in game locations",
			callback: commandMapNext,
		},
		"mapp": {
			name:"map prev",
			description:"Returns prev list of in game locations",
			callback: commandMapPrev,
		},
		"explore": {
			name:"explore",
			description:"Returns list of pokemon for specific location",
			callback: commandExplore,
		},
	}	
}

