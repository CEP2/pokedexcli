package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	result := text
	result = strings.Trim(text, " ")
	result = strings.ToLower(result)

	fieldsResult := strings.Fields(result)

	return fieldsResult
}

func startRepl(cfg *config) error {
	scanner := bufio.NewScanner(os.Stdin)
	err := scanner.Err()
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	//main assessment loop
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())[0]
		if cmd, ok := commandMap[input]; ok {
			cmd.callback(cfg)
		} else {
			fmt.Printf("Unknown command\n")
		}

	}

	return nil
}
