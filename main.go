package main

import (
	"time"

	"github.com/cep2/pokedexcli/internal/pokecache"
)

func main() {

	//build cache
	duration := 5 * time.Second
	cfg := &config{
		cmdRegistry: commandsList,
		cache:       pokecache.NewCache(duration),
	}

	startRepl(cfg)

}
