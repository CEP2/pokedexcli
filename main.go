package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// fmt.Println("Hello, World!")
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
		fmt.Printf("Your command was: %s\n", input)
	}
}
