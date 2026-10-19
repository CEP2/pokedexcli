package main

type config struct {
	state       string
	cmdRegistry map[string]cliCommand
}
