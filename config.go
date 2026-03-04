package main

import (
	"github.com/stampersoftware/pokedex/internal"
)

type Config struct {
	location *LocationConfig
	cache *internal.Cache
	client *internal.Client
	pokedex *Pokedex
}

func initConfig() (Config, error) {
	
	cache, err := internal.NewCache(20)
	client := internal.NewClient()
	location := initLocationConfig()
	pokedex := CreatePokedex()

	if err != nil {
		return Config{}, err
	}
	
	config := Config{
		location: location,
		cache: cache,
		client: client,
		pokedex : pokedex,
	}

	return config, nil
}
