package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandExplore(cfg *config, location ...string) error {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	//receive location
	if len(location) == 0 {
		return fmt.Errorf("location not provided.\t\nexplore <location-area Name or ID>.")
	}
	loc := location[0]
	locAPI := baseURL + loc
	//check cache if API url hasn't been called yet.
	data, cached := cfg.cache.Get(locAPI)
	// if not call API
	if !cached {
		res, err := http.Get(locAPI)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("response failed: %v\n\t'%s' is not a location.", res.StatusCode, loc)
		}
		// read data from response
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		//add to cache
		cfg.cache.Add(locAPI, data)
	}

	//process data
	var locationArea LocationAreaResponse
	err := json.Unmarshal(data, &locationArea)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring: %s\n", loc)
	for _, pkmn := range locationArea.PokemonEncounters {
		fmt.Printf("\t- %s\n", pkmn.Pokemon.Name)
	}
	//store API in cache
	cfg.cache.Add(locAPI, data)

	return nil
}
