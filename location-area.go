package main
import (
	"encoding/json"
	"net/http"
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

type LocationArea struct {
	
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

type LocationAreaResponse struct {
	Pokemon_Encounters []struct{
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
	
}

type Pokemon struct {
	Name string `json:"name"`
	URL string `json:"url"`
}


func GetLocation(location string, config *Config) (LocationAreaResponse, error) {

	data := LocationAreaResponse{}
	url := locationUrl + "/" + location
	if val, found := config.cache.Get(url); found == true {
		if err := json.Unmarshal(val, &data); err != nil {
			return LocationAreaResponse{}, err
		}
	} else {

	
		req, err := http.NewRequest("GET", locationUrl + "/" + location, nil)
	
		if err != nil {
			return LocationAreaResponse{}, nil
		}

		res, err := config.client.Client.Do(req)
		if err != nil {
			return LocationAreaResponse{}, nil
		}

		defer res.Body.Close()
	
		decoder := json.NewDecoder(res.Body)
	
		err = decoder.Decode(&data)
	
		if err != nil {
			return LocationAreaResponse{}, nil
		}
	}

	return data, nil
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

