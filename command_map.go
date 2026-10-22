package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(cfg *config, args ...string) error {
	mapAPI := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != nil {
		mapAPI = *cfg.next
	}

	//check cache for existing data
	data, cached := cfg.cache.Get(mapAPI)
	if !cached {
		// call api for pokemon location.
		res, err := http.Get(mapAPI)
		if err != nil {
			return err
		}
		// parse json get individual lines.
		defer res.Body.Close()
		if res.StatusCode > 299 {
			fmt.Errorf("response failed: %v", res.StatusCode)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.cache.Add(mapAPI, data)
	}

	var locationArea LocationAreaResponse
	err := json.Unmarshal(data, &locationArea)
	if err != nil {
		return err
	}

	for _, loc := range locationArea.Results {
		fmt.Println(loc.Name)
	}

	//store in cache
	cfg.cache.Add(mapAPI, data)

	//update next call
	cfg.next = locationArea.Next
	cfg.previous = locationArea.Previous

	return nil
}
