package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(c *Config, name []string) error {
	
	if len(name) == 0 {
		return errors.New("must have pokemon name")
	}
	pokemon, err := GetPokemon(name[0], c)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a pokeball at %s\n", pokemon.Name)
	
	if attemptCatch(pokemon.Base_Experience) {
		c.pokedex.AddToPokedex(pokemon)
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s got away!\n", pokemon.Name)
	}

	return nil
}


func attemptCatch(baseExperience int) bool {
	return rand.Intn(baseExperience) > baseExperience / 2 
}
