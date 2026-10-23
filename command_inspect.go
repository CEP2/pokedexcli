package main

import "fmt"

func commandInspect(cfg *config, pokemon ...string) error {
	//grab pokemon
	if len(pokemon) == 0 {
		return fmt.Errorf("pokemon not provided.\t\nusage: inspect <location-area Name or ID>.")
	}

	name := pokemon[0]
	//validate pokemon has been caught

	if pkmn, ok := cfg.pokedex[name]; !ok {
		fmt.Printf("%s is not an entry in your pokedex. Go catch some more pokemon!\n", name)
	} else {
		// fmt.Print(pkmn.Stats)

		// get size...
		h := pkmn.Height
		w := pkmn.Weight
		// get stats
		stats := pkmn.Stats
		hp := stats[0].BaseStat
		atk := stats[1].BaseStat
		def := stats[2].BaseStat
		spAtk := stats[3].BaseStat
		spDef := stats[4].BaseStat
		spd := stats[5].BaseStat

		fmt.Printf(`Name: %s
Height: %d
Weight: %d
Stats:
	- HP: 		%d
	- ATK: 		%d
	- DEF:		%d
	- Sp.ATK	%d
	- Sp.DEF	%d
	- SPD		%d
Types:
`, name, h, w, hp, atk, def, spAtk, spDef, spd)
		for _, pkmnType := range pkmn.Types {
			fmt.Printf("\t- %s\n", pkmnType.Type.Name)
		}
	}
	return nil
}
