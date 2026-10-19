package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	err := scanner.Err()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	//main assessment loop
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())[0]
		if cmd, ok := commandMap[input]; ok {
			cmd.callback()
		} else {
			fmt.Printf("Unknown command\n")
		}

	}
}
