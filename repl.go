package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)



func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	prompt := "Pokedex > "
	var userInput []string
	for {
		fmt.Print(prompt)
		if !scanner.Scan() {
			fmt.Println("Error reading input:", scanner.Err())
			return
		}


		userInput = cleanInput(scanner.Text())

		if len(userInput) == 0 {
			continue
		}
		commandName := userInput[0]
		command, exists := commands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	var result []string
	result = strings.Fields(strings.ToLower(strings.TrimSpace(text)))
	return result
}
