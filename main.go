package main
import (
	"os"
	"fmt"
)

func main() {
		
	config, err := initConfig()
	
	if err != nil {
		fmt.Printf("Error with config:\n-%v\n", err)
		os.Exit(1)
	}
	
	startRepl(&config)	
	
}

