package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {

	if len(cfg.pokedex) == 0 {
		fmt.Printf("Your pokedex is empty! Catch some pokemon!\n")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, pkmn := range cfg.pokedex {
		fmt.Printf("\t- %s\n", pkmn.Name)
	}

	return nil

}
