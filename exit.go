package main

import (
	"os"
	"fmt"
)
func commandExit(c *Config, params []string) error {
	fmt.Println("Goodbye")
	os.Exit(0)
	return nil
}

