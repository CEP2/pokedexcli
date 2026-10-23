package main

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, args ...string) error
}

var commandsList = map[string]cliCommand{
	"catch": {
		name:        "catch",
		description: "attempts to catch a pokemon. usage: 'catch <pokemon>'",
		callback:    commandCatch,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
	"explore": {
		name:        "explore",
		description: "explores a location. usage: 'explore <map location>'",
		callback:    commandExplore,
	},
	"help": {
		name:        "help",
		description: "Prints command list",
		callback:    commandHelp,
	},
	"inspect": {
		name:        "inspect",
		description: "shows a caught pokemon's stats. usage: 'inspect <pokemon>'",
		callback:    commandInspect,
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
	"pokedex": {
		name:        "pokedex",
		description: "displays caught pokemon",
		callback:    commandPokedex,
	},
}
