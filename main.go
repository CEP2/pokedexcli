package main

func main() {

	cfg := &config{cmdRegistry: commandMap}

	startRepl(cfg)

}
