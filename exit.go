package main

import (
	"os"
	"fmt"
)
func commandExit(c *Config) error {
	fmt.Println("Goodbye")
	os.Exit(0)
	return nil
}

