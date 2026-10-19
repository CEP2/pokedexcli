package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commandMap = map[string]cliCommand{
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
}
