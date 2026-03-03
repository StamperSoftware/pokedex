package main

import (
	"github.com/stampersoftware/pokedex/internal"
)

type Config struct {
	location *LocationConfig
	cache *internal.Cache
	client *internal.Client
}

func initConfig() (Config, error) {
	
	cache, err := internal.NewCache(20)
	client := internal.NewClient()
	location := initLocationConfig()

	if err != nil {
		return Config{}, err
	}
	
	config := Config{
		location: location,
		cache: cache,
		client: client,
	}

	return config, nil
}
