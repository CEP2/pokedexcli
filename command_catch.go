package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func commandCatch(cfg *config, pokemon ...string) error {

	//Endpoint
	//https://pokeapi.co/api/v2/pokemon/{id or name}/

	baseURL := "https://pokeapi.co/api/v2/pokemon/"

	if len(pokemon) == 0 {
		return fmt.Errorf("pokemon not provided.\t\nusage: catch <location-area Name or ID>.")
	}

	pkmn := pokemon[0]
	pkmnAPI := baseURL + pkmn

	//check cache
	data, cached := cfg.cache.Get(pkmnAPI)
	if !cached {
		res, err := http.Get(pkmnAPI)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		// check if response valid
		if res.StatusCode > 299 {
			return fmt.Errorf("response failed: %v\n\t'%s' is not a pokemon.", res.StatusCode, pkmn)
		}

		//read data from response
		data, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		//add to cache on new query
		cfg.cache.Add(pkmnAPI, data)
	}

	var pokemonData Pokemon
	err := json.Unmarshal(data, &pokemonData)
	if err != nil {
		return err
	}

	name := pkmn
	baseExp := pokemonData.BaseExperience

	//Catching
	//Divide baseExp by this, then 1/X is chance for catching.
	//If baseexp is less than 20, make chance 1/2
	divChance := 10
	chanceCatch := 2
	if baseExp > 20 {
		chanceCatch = baseExp / divChance
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	catchAttempt := rand.Intn(chanceCatch)

	if catchAttempt != 0 {
		fmt.Printf("\tMaybe next time, %s broke free!\n", name)
	} else {
		fmt.Printf("\tCongratulations, %s was caught! Updating pokedex...\n", name)
		cfg.pokedex[name] = pokemonData
	}
	return nil
}
