package main
import (
	"encoding/json"
	"net/http"
	"fmt"
)

const locationUrl = "https://pokeapi.co/api/v2/location-area"

type LocationConfig struct {
	next string
	prev string
}

type Location struct {
	URL string `json:"url"`
	Name string `json:"name"`
}

func initLocationConfig() *LocationConfig {
	return &LocationConfig {
		next : locationUrl,
		prev : "",
	}
}

type LocationResponse struct {
	Count int `json:"count"`
	Results []Location `json:"results"`
	Next string `json:"next"`
	Previous string `json:"previous"`
}

func GetLocations(url string, config *Config) (LocationResponse, error) {
	
	data := LocationResponse{}

	if val, found := config.cache.Get(url); found == true {
		if err := json.Unmarshal(val, &data); err != nil {
			return LocationResponse{}, err
		}

	} else {
 		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationResponse{}, err
		}	
		
		res, err := config.client.Client.Do(req)
		if err != nil {
			return LocationResponse{}, err
		}
		defer res.Body.Close()
		decoder := json.NewDecoder(res.Body)
		err = decoder.Decode(&data)
	
		if err != nil {
			return LocationResponse{}, err
		}
		parsedData, err := json.Marshal(data)
		if err != nil {
			return LocationResponse{}, err
		}
		config.cache.Add(url, parsedData)
	}	

	config.location.next = data.Next
	config.location.prev = data.Previous
	return data, nil
}

