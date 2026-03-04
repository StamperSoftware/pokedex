package main

import (
	"errors"
	"fmt"
)

func commandExplore(c *Config, params []string) error {
	if len(params) == 0 {
		return errors.New("must have location name")
	}
	location, err := GetLocation(params[0], c)
	if err != nil {
		return err
	}
	
	for _, pokemonDto := range location.Pokemon_Encounters {
		fmt.Println(pokemonDto.Pokemon.Name)
	}

	return nil
}
