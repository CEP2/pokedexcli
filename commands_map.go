package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

var commandsList = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"help": {
		name:        "help",
		description: "Prints command list",
		callback:    commandHelp,
	},
	"map": {
		name:        "map",
		description: "displays Pokemon map locations. repeated calls display next 20 locations",
		callback:    commandMap,
	},
	"mapb": {
		name:        "mapb",
		description: "displays previous 20 Pokemon map locations.",
		callback:    commandMapb,
	},
	"explore": {
		name:        "explore",
		description: "explores a location. usage: 'explore <map location>'",
		callback:    commandExplore,
	},
	"catch": {
		name:        "catch",
		description: "attempts to catch a pokemon. usage: 'catch <pokemon>'",
		callback:    commandCatch,
	},
}
