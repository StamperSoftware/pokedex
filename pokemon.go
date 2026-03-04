package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

const pokemonUrl = "https://pokeapi.co/api/v2/pokemon"

type Pokemon struct {
	Name string `json:"name"`
	Base_Experience int `json:"base_experience"`

}

func GetPokemon(name string, config *Config) (Pokemon, error) {
	
	data := Pokemon{}

	url := pokemonUrl + "/" + name
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := config.client.Client.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()
	
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&data)
	if err != nil {
		return Pokemon{}, err
	}

	return data, nil
}

func(p *Pokemon) ListDetails() {
	fmt.Printf("Name: %s\nBase Experience:%d\n", p.Name, p.Base_Experience)
}
