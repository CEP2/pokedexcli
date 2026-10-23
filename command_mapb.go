package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapb(cfg *config, args ...string) error {
	var mapAPI string

	if cfg.previous != nil {
		mapAPI = *cfg.previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}
	// fmt.Printf(*cfg.previous)
	// check cache before request
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
			return fmt.Errorf("response failed: %v", res.StatusCode)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.cache.Add(mapAPI, data)
	}

	var locationAreas LocationAreasResponse
	err := json.Unmarshal(data, &locationAreas)
	if err != nil {
		return err
	}

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}

	//update next call
	cfg.next = locationAreas.Next
	cfg.previous = locationAreas.Previous

	return nil
}
