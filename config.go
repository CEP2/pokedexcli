package main

import "github.com/cep2/pokedexcli/internal/pokecache"

type config struct {
	cmdRegistry map[string]cliCommand
	pokedex     map[string]Pokemon
	cache       *pokecache.Cache
	next        *string
	previous    *string
}
