package main

import (
	"fmt"
	"errors"
)

type Pokedex struct {
	pokemon map[string]Pokemon
}

func CreatePokedex() *Pokedex {
	return &Pokedex{
		pokemon : make(map[string]Pokemon),
	}
}

func(p *Pokedex) AddToPokedex(pokemon Pokemon) {
	p.pokemon[pokemon.Name] = pokemon
}

func commandPokedex(config *Config, _ []string) error {
	for _, pokemon := range config.pokedex.pokemon {
		fmt.Printf("Name: %s, experience:%d\n",pokemon.Name, pokemon.Base_Experience)
	}
	return nil
}

func commandInspect(config *Config, names []string) error {
	if len(names) == 0 {
		return errors.New("must have a name")
	}
	if pokemon, wasFound := config.pokedex.pokemon[names[0]]; wasFound == true {
		pokemon.ListDetails()
	} else {
		fmt.Printf("%s has not been caught\n", names[0])
	}
	return nil
}
