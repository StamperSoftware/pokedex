package main

import (
	"fmt"
	"errors"
)

func commandMapNext(config *Config, params []string) error {
	
	locations, err := GetLocations(config.location.next, config)
	
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	
	return nil
}

func commandMapPrev(config *Config, params []string) error {
	
	if config.location.prev == "" {
		return errors.New("on first page")
	}
	
	locations, err := GetLocations(config.location.prev, config)
	
	if err != nil {
		fmt.Println(err)
		return err
	}
	
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	
	return nil
}


